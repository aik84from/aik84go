/*

*- ** -*- ---** ****-

Version: 0.0.1

Author: Kalinin Alexandr

Email: aik84from@gmail.com

Github: https://github.com/aik84from

8b7521f0bc0057306aeee70bf6603feb

*/

package aik84go

import (
    "math"
    "math/rand"
)

func Factorial(n uint64) uint64 {
    if n == 1 {
        return 1
    }
    return n * Factorial(n - 1)
}

func KineticEnergy(mass, velocity float64) float64 {
    return (mass * math.Pow(velocity, 2)) * 0.5
}

func Watts(volts, amperes float64) float64 {
    return volts * amperes
}

func HorsepowerInternalCombustionEngine(torque, RPM float64) float64 {
    return (torque * RPM) / 5252.0
}

func SphereVolume(radius float64) float64 {
    return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func CylinderVolume(radius, height float64) float64 {
    return math.Pi * math.Pow(radius, 2) * height
}

func ConeVolume(radius, height float64) float64 {
    return math.Pi * math.Pow(radius, 2) * (height / 3.0)
}

func FreeFallTimeEarth(height float64) float64 {
    return math.Sqrt(2.0 * height / GravitationalAccelerationEarth())
}

func FreeFallVelocityEarth(height float64) float64 {
    return math.Sqrt(2.0 * GravitationalAccelerationEarth() * height)
}

func HydrostaticPressureWater(height float64) float64 {
    return (WaterDensity() * GravitationalAccelerationEarth() * height) + SeaLevelStandardAtmosphericPressureEarth()
}

func SimpleLinearRegression(intercept, slope, x float64) float64 {
    return intercept + (slope * x)
}

func Sigmoid(x float64) float64 {
    return 1.0 / (1.0 + math.Exp(-x))
}

func RoundFloat(x, n float64) float64 {
    return math.Round(x * n) / n
}

func Sgn(x float64) int {
    switch {
        case x < 0.0:
            return -1
        case x > 0.0:
            return 1
    }
    return 0
}

func Between(x, from, to float64) bool {
    return x >= from && x <= to
}

func RandInt(min, max int) int {
    return rand.Intn(max - min) + min
}

func RandFloat(min, max float64) float64 {
    return min + rand.Float64() * (max - min)
}

func Probability(favourable, total float64) float64 {
    return favourable / total
}

func EuclideanDistance(a, b []float64) float64 {
    if len(a) != len(b) {
        panic("len(a) != len(b)")
    }
    var result float64 = 0.0
    for i := 0; i < len(a); i++ {
        result += math.Pow(a[i] - b[i], 2)
    }
    return math.Sqrt(result)
}

func ManhattanDistance(a, b []float64) float64 {
    if len(a) != len(b) {
        panic("len(a) != len(b)")
    }
    var result float64 = 0.0
    for i := 0; i < len(a); i++ {
        result += math.Abs(a[i] - b[i])
    }
    return result
}

func SpeedOfLightInVacuum() float64 {
    return 299792458.0
}

func GravitationalAccelerationEarth() float64 {
    return 9.80665
}

func SeaLevelStandardAtmosphericPressureEarth() float64 {
    return 101325.0
}

func WaterDensity() float64 {
    return 1000.0
}

func MorseCodeNumbers() map[string]string {
    morseCode := make(map[string]string)
    
    morseCode["-----"] = "0"
    morseCode["*----"] = "1"
    morseCode["**---"] = "2"
    morseCode["***--"] = "3"
    morseCode["****-"] = "4"
    morseCode["*****"] = "5"
    morseCode["-****"] = "6"
    morseCode["--***"] = "7"
    morseCode["---**"] = "8"
    morseCode["----*"] = "9"
    
    morseCode["0"] = "-----"
    morseCode["1"] = "*----"
    morseCode["2"] = "**---"
    morseCode["3"] = "***--"
    morseCode["4"] = "****-"
    morseCode["5"] = "*****"
    morseCode["6"] = "-****"
    morseCode["7"] = "--***"
    morseCode["8"] = "---**"
    morseCode["9"] = "----*"
    
    return morseCode
}
