package goods

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFollow(t *testing.T) {
	arr := [][2]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{3, 1},
		{3, 2},
		{3, 3},
	}

	f := NewFollow()
	Convey("follow select", t, func() {
		Convey("follow", func() {
			for _, v := range arr {
				So(f.Follow(v[0], v[1]), ShouldBeTrue)
				So(f.Check(v[0], v[1]), ShouldBeTrue)
			}
		})
		Convey("followings", func() {
			So(f.UserFollowings(1, 1, 10), ShouldNotBeNil)
			So(f.UserFollowings(2, 1, 10), ShouldBeNil)
		})
		Convey("unfollow", func() {
			for _, v := range arr {
				So(f.Unfollow(v[0], v[1]), ShouldBeTrue)
				So(f.Check(v[0], v[1]), ShouldBeFalse)
			}
		})
	})

}
