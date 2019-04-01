package generation

import (
	"reflect"
	"testing"
)

func TestNoisyWorkloadGenerator_GenerateWorkload(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		n    *NoisySin
		args args
		want []float64
	}{
		{
			"High noise",
			&NoisySin{
				Amplitude: 120,
				Noise:     60,
				Period:    60,
			},
			args{
				length: 120,
			},
			[]float64{336.6818099599512, 304.3057086145077, 295.09395584967484, 219.90584773540272, 251.0385425349977, 304.06365325562535, 260.0114234957638, 291.2782899897749, 334.63676046564296, 320.60914069958653, 406.29376495609483, 340.1251697463514, 340.2193815915308, 339.1419378115181, 243.0858165970983, 312.1495016774793, 247.69199273915007, 278.00828728002045, 329.90988448119447, 191.63993346610562, 345.11650816596557, 248.47976226067829, 202.8964285714095, 345.3410129849017, 267.80376790479295, 237.4654644017046, 306.0621085316434, 245.80619196855645, 177.06253538142784, 226.13329797699356, 281.0315773633437, 240.66711159352764, 264.07324276786426, 212.0889454774939, 301.38438461154817, 278.5425285129831, 179.94097344012647, 284.24175933364995, 189.5233125547872, 221.61768540169118, 240.42372469419146, 178.2165721864046, 215.7339929609601, 175.3021513019661, 178.63046163535444, 285.6108437723026, 319.5863889162888, 203.58260039989463, 255.84914543326363, 202.27138446076708, 231.61324324835113, 201.72005786431004, 216.73278933573548, 254.30535701966414, 214.39406202896328, 172.67498142526463, 148.62922990461067, 214.97308370203064, 228.3887548812695, 173.23134549271077, 245.68943267157087, 104.05723259500368, 144.43685959018936, 121.43713077346908, 211.26055172552785, 207.06394321105319, 146.68075231702213, 122.20315507934582, 130.38691903670528, 151.8633546995093, 86.25219707979923, 177.32463731776141, 197.59318264264834, 88.7762258556206, 77.31183145939335, 165.7765051260893, 211.98003887570945, 48.91919623160378, 117.06212444693712, 23.18212569792179, 121.82945300950635, 90.94454773535165, 171.60418606483202, 56.87688902413155, 148.28986966985522, 177.38854184675, 221.4673865917412, 84.05613566044184, 71.22322517833855, 76.58375612036905, -14.038629384259778, 109.74592591293657, 70.10767613590008, 128.57195219257926, 156.0844771924046, 118.16589179913613, 153.29015825123886, 99.3389839921977, 17.629069708587494, 124.3299924335776, -11.07238407435625, 72.34267601879375, 61.79070779912179, 80.76466902209705, 57.835755693454274, 79.81933638373293, 8.31152931869822, 118.08388998765523, 0.9821901494387077, 116.31635031846703, 108.70394654081046, 44.9446755009757, 175.982945333643, 123.51334552378437, 151.54421118952592, -27.77185595261774, 30.590769307108143, -35.404296396861554, -14.770316723913382, 149.73728446297253},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.GenerateWorkload(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoisySin.GenerateWorkload() = %v, want %v", got, tt.want)
			}
		})
	}
}