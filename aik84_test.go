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
    "testing"
)

func TestWatts(t *testing.T) {
    result := Watts(10, 2)
    if result != 20 {
        t.Fatal(result)
    }
}

func TestFactorial(t *testing.T) {
    result := Factorial(8)
    if result != 40320 {
        t.Fatal(result)
    }
}

func TestKineticEnergy(t *testing.T) {
    result := KineticEnergy(5, 10)
    if result != 250 {
        t.Fatal(result)
    }
}

func TestSigmoid(t *testing.T) {
    result := Sigmoid(5.1532)
    if result != 0.9942523501596811973488 {
        t.Fatal(result)
    }
}

func TestRandInt(t *testing.T) {
    result := RandInt(10, 20)
    if result < 10 || result > 20 {
        t.Fatal(result)
    }
}

func TestRandFloat(t *testing.T) {
    result := RandFloat(10.0, 20.0)
    if result < 10.0 || result > 20.0 {
        t.Fatal(result)
    }
}

func TestSgn(t *testing.T) {
    if Sgn(8) != 1 {
        t.Fatal(`Sigmoid (1)`)
    }
    if Sgn(0) != 0 {
        t.Fatal(`Sigmoid (0)`)
    }
    if Sgn(-3) != -1 {
        t.Fatal(`Sigmoid (-1)`)
    }
}

func TestBetween(t *testing.T) {
    if Between(10, 8, 22) != true {
        t.Fatal(`Between`)
    }
}

func TestProbability(t *testing.T) {
    result := Probability(5, 10)
    if result != 0.5 {
        t.Fatal(result)
    }
}

func TestRoundFloat(t *testing.T) {
    result := RoundFloat(9.547854135, 1000)
    if result != 9.548 {
        t.Fatal(result)
    }
}

func TestSimpleLinearRegression(t *testing.T) {
    result := SimpleLinearRegression(2.5, 2.0, 4.0)
    if result != 10.5 {
        t.Fatal(result)
    }
}

func TestEuclideanDistance(t *testing.T) {
    result := EuclideanDistance([]float64{7, 22}, []float64{49, 73})
    if result != 66.06814663663572 {
        t.Fatal(result)
    }
}

func TestManhattanDistance(t *testing.T) {
    result := ManhattanDistance([]float64{7, 22}, []float64{49, 73})
    if result != 93 {
        t.Fatal(result)
    }
}

func TestSphereVolume(t *testing.T) {
    result := SphereVolume(10)
    if result != 4188.790204786392 {
        t.Fatal(result)
    }
}

func TestCylinderVolume(t *testing.T) {
    result := CylinderVolume(10, 5)
    if result != 1570.7963267948967 {
        t.Fatal(result)
    }
}

func TestConeVolume(t *testing.T) {
    result := ConeVolume(10, 5)
    if result != 523.5987755982989 {
        t.Fatal(result)
    }
}

func TestHorsepowerInternalCombustionEngine(t *testing.T) {
    result := HorsepowerInternalCombustionEngine(500, 2000)
    if result != 190.4036557501904 {
        t.Fatal(result)
    }
}

func TestHydrostaticPressureWater(t *testing.T) {
    result := HydrostaticPressureWater(17)
    if result != 268038.05 {
        t.Fatal(result)
    }
}

func TestFreeFallTimeEarth(t *testing.T) {
    result := FreeFallTimeEarth(150)
    if result != 5.530957095235674 {
        t.Fatal(result)
    }
}

func TestFreeFallVelocityEarth(t *testing.T) {
    result := FreeFallVelocityEarth(150)
    if result != 54.24016039799292 {
        t.Fatal(result)
    }
}

func TestSpeedOfLightInVacuum(t *testing.T) {
    result := SpeedOfLightInVacuum()
    if result != 299792458 {
        t.Fatal(result)
    }
}

func TestGravitationalAccelerationEarth(t *testing.T) {
    result := GravitationalAccelerationEarth()
    if result != 9.80665 {
        t.Fatal(result)
    }
}

func TestWaterDensity(t *testing.T) {
    result := WaterDensity()
    if result != 1000 {
        t.Fatal(result)
    }
}

func TestSeaLevelStandardAtmosphericPressureEarth(t *testing.T) {
    result := SeaLevelStandardAtmosphericPressureEarth()
    if result != 101325.0 {
        t.Fatal(result)
    }
}

func TestMorseCodeNumbers(t *testing.T) {
    result := MorseCodeNumbers()
    
    if result["8"] != "---**" {
        t.Fatal(result["8"])
    }
    if result["***--"] != "3" {
        t.Fatal(result["***--"])
    }
}

