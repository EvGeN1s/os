package mimization

import (
	"fmt"
	"os/lw2/convert"
	"os/lw2/model"
)

func Mealy(a model.Mealy) model.Mealy {
	eg := convert.MealyToEqualityGroup(a)
	fmt.Println(eg)
	eg = EqualityGroup(eg)
	fmt.Println(eg, len(eg.StateToGroup))
	return convert.EqualityGroupToMealy(a, eg)
}
