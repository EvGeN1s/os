package mimization

import (
	"os/lw2/convert"
	"os/lw2/model"
)

func Moore(a model.Moore) model.Moore {
	eg := convert.MooreToEqualityGroup(a)
	eg = EqualityGroup(eg)
	return convert.EqualityGroupToMoore(a, eg)
}
