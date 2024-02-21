package problemdetail

func BadRequestProblemDetail(message string) ProblemDetail {
	return ProblemDetail{
		Type:    ProblemDetailRootSchema + "BadRequest",
		Title:   "Invalid data in the request body",
		Detail:  message,
		Context: nil,
	}
}

func NotFoundProblemDetail(message string) ProblemDetail {
	return ProblemDetail{
		Type:    ProblemDetailRootSchema + "NotFound",
		Title:   "Resource Not Found",
		Detail:  message,
		Context: nil,
	}
}

func UnauthorizedProblemDetail(message string) ProblemDetail {
	return ProblemDetail{
		Type:    ProblemDetailRootSchema + "Unauthorized",
		Title:   "Unauthorized to access resource or perform operation",
		Detail:  message,
		Context: nil,
	}
}

func PaymentRequiredProblemDetail(message string) ProblemDetail {
	return ProblemDetail{
		Type:    ProblemDetailRootSchema + "PaymentRequired",
		Title:   "Payment is required to perform the operation....",
		Detail:  message,
		Context: nil,
	}
}

func ServerErrorProblemDetail(message string) ProblemDetail {
	return ProblemDetail{
		Type:    ProblemDetailRootSchema + "InternalServerError",
		Title:   "Internal Server Error occured while processing the request",
		Detail:  message,
		Context: nil,
	}
}
