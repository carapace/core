package ark

func WithDefault(s *Service) {
	s.validator = newValidator()
	s.validator.txmiddleware = append(
		s.validator.txmiddleware,

		amountIsSet,
		feeIsSet,
		typeIsSet,
	)
}
