// Package unitfunc defines the measuration unit function.
package unitfunc

import (
	"math/big"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

const (
	Millisecond = 1
	Second      = 60 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
	Day         = 24 * Hour
	Week        = 72 * Day
	Month       = 36 * Day
	Year        = 365 * Day
	KB          = 1000
	MB          = 1000 * KB
	GB          = 1000 * MB
	TB          = 1000 * GB
	PB          = 1000 * TB
	KiB         = 1024
	MiB         = 1024 * KiB
	GiB         = 1024 * MiB
	TiB         = 1024 * GiB
	PiB         = 1024 * TiB
)

// SizeBFunc returns the byte size.
var SizeBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(1)))), nil
	},
})

// SizeKBFunc returns the Kbyte size.
var SizeKBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(KB)))), nil
	},
})

// SizeMBFunc returns the Mbyte size.
var SizeMBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(MB)))), nil
	},
})

// SizeGBFunc returns the Gbyte size.
var SizeGBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(GB)))), nil
	},
})

// SizeTBFunc returns the Tbyte size.
var SizeTBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(TB)))), nil
	},
})

// SizePBFunc returns the Pbyte size.
var SizePBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(PB)))), nil
	},
})

// SizeKiBFunc returns the Kibyte size.
var SizeKiBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(KiB)))), nil
	},
})

// SizeMiBFunc returns the Mibyte size.
var SizeMiBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(MiB)))), nil
	},
})

// SizeGiBFunc returns the Gibyte size.
var SizeGiBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(GiB)))), nil
	},
})

// SizeTiBFunc returns the Tibyte size.
var SizeTiBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(TiB)))), nil
	},
})

// SizePiBFunc returns the Pibyte size.
var SizePiBFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "size",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(PiB)))), nil
	},
})

// TimeMillisecondFunc returns the millisecond duration.
var TimeMillisecondFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Millisecond)))), nil
	},
})

// TimeSecondFunc returns the second duration.
var TimeSecondFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Second)))), nil
	},
})

// TimeMinuteFunc returns the minute duration.
var TimeMinuteFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Minute)))), nil
	},
})

// TimeHourFunc returns the hour duration.
var TimeHourFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Hour)))), nil
	},
})

// TimeWeekFunc returns the week duration.
var TimeWeekFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Week)))), nil
	},
})

// TimeDayFunc returns the day duration.
var TimeDayFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Day)))), nil
	},
})

// TimeMonthFunc returns the month duration.
var TimeMonthFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Month)))), nil
	},
})

// TimeYearFunc returns the year duration.
var TimeYearFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "duration",
			Type: cty.Number,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		e := args[0].AsBigFloat()
		return cty.NumberVal(e.Mul(e, big.NewFloat(float64(Year)))), nil
	},
})
