package helpers

import "Solver_equation_API/models"

func Convertor(coef *models.Coef, answer *models.Answer) {
	answer.A = coef.A
	answer.B = coef.B
	answer.C = coef.C
	answer.Roots = 0
}
