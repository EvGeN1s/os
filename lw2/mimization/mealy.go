package mimization

import (
	"os/lw2/convert"
	"os/lw2/model"
)

func Mealy(a model.Mealy) model.Mealy {
	eg := convert.MealyToEqualityGroup(a)
	eg = EqualityGroup(eg)
	return convert.EqualityGroupToMealy(a, eg)
}
