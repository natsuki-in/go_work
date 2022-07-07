package main

import "math"

func calcBMI(height float64, weight float64) float64 {
  // 身長をcm換算からm換算に変更
  height = height / 100
  // BMIの計算結果を返す
  return weight / math.Pow(height, 2)
}