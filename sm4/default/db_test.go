package _default

import (
	"fmt"
	"github.com/alserov/gb/sm4/default/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDbImpl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := cachemock.NewMockCache(ctrl)
	c.EXPECT().
		Get(gomock.Eq("key")).
		Return(nil, false).
		Times(1)
	c.EXPECT().
		Set(gomock.Eq("key"), gomock.Eq([]byte("val"))).
		Times(1)

	db := newDbImpl(c)

	db.Insert("key", []byte("val"))

	val, ok := db.Get("key")
	require.True(t, ok)
	require.Equal(t, []byte("val"), val)

}

func TestDbImpl_ScheduleInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := cachemock.NewMockCache(ctrl)
	c.EXPECT().
		GetWriteCache(gomock.Eq(true)).
		Return(map[string][]byte{"key": []byte("val")}).
		AnyTimes()
	c.EXPECT().Get(gomock.Any()).Return(nil, false).Times(2)
	c.EXPECT().Set("key", []byte("val")).Times(1)

	db := newDbImpl(c)
	go db.ScheduleInsert(time.NewTicker(time.Second))

	v, ok := db.Get("key")
	require.False(t, ok)
	require.Empty(t, v)

	time.Sleep(time.Second * 2)

	v, ok = db.Get("key")
	require.True(t, ok)
	require.Equal(t, []byte("val"), v)
}

func TestDbImpl_LimitInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := cachemock.NewMockCache(ctrl)
	c.EXPECT().
		GetWriteCache(gomock.Any()).
		Return(map[string][]byte{"key0": []byte("val0"), "key1": []byte("val1"), "key2": []byte("val2")}).
		AnyTimes()
	c.EXPECT().Get(gomock.Any()).Return(nil, false).AnyTimes()
	c.EXPECT().SetWriteCache(gomock.Any(), gomock.Any()).AnyTimes()
	c.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()

	db := newDbImpl(c)
	go db.LimitInsert(3)

	for i := 0; i < 3; i++ {
		c.SetWriteCache(fmt.Sprintf("key%d", i), []byte(fmt.Sprintf("val%d", i)))
		v, ok := db.Get(fmt.Sprintf("key%d", i))
		require.False(t, ok)
		require.Empty(t, v)

		if i == 2 {
			time.Sleep(time.Second * 2)
			v, ok = db.Get(fmt.Sprintf("key%d", i))
			require.True(t, ok)
			require.Equal(t, fmt.Sprintf("val%d", i), string(v))
		}
	}
}
