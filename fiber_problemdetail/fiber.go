package problemdetail

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NotAuthorized(ctx *fiber.Ctx, message string) error {
	return ctx.Status(http.StatusUnauthorized).JSON(UnauthorizedProblemDetail(message))
}

func ServerError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(http.StatusInternalServerError).JSON(ServerErrorProblemDetail(message))
}

func BadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(http.StatusBadRequest).JSON(BadRequestProblemDetail(message))
}

func NotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(http.StatusNotFound).JSON(NotFoundProblemDetail(message))
}

func ValidationErrors(ctx *fiber.Ctx, message string, validationErrors any) error {
	return ctx.Status(http.StatusBadRequest).JSON(ProblemDetail{
		Type:    ProblemDetailRootSchema + "ValidationErrors",
		Title:   "Invalid data in the request body",
		Detail:  message,
		Context: validationErrors,
	})
}
