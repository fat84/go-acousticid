package acousticid
import (
	"os/exec"
	"strings"
	"strconv"
)

type Fingerprint struct {
	fingerprint string
	duration float64
}

func (f *Fingerprint) Get(file string) {
	out, err := exec.Command("./fpcalc", file).Output()
	if err != nil {
		panic(err)
	}
	outstrs := strings.Split(string(out), "\n")

	for _, s := range outstrs  {
		if strings.Index(s, "DURATION=") == 0 {
			ds := strings.Split(s, "=")[1]
			f.duration, _ = strconv.ParseFloat(ds, 64)
		} else if strings.Index(s, "FINGERPRINT=") == 0 {
			f.fingerprint = strings.Split(s, "=")[1]
		}
	}
}
