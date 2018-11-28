package dispatcher

// func (s *Service) TransactionService(ctx context.Context, trs *v0.Transaction) (res *v0.Transaction, err error) {
// 	tx := core.TXFromContext(ctx)
// 	identity, err := s.store.Sets.Identity.Get(ctx, tx, trs.Identity)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	signingUser, err := s.store.Users.Get(ctx, tx, trs.Witness.Signatures[0].GetPrimaryPublicKey())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	policies, err := permissions.PoliciesFromProto(identity.Policies, []string{identity.Name})
//
// 	err = s.perm.DoPoliciesAllow(&ladon.Request{
// 		Subject:  signingUser.Name,
// 		Resource: identity.Name,
// 		Action:   v0.Action_Use.String(),
// 		Context: map[string]interface{}{
// 			v0.ConditionNames_AuthLevelGTE.String(): signingUser,
// 			v0.ConditionNames_UsersOwns.String():    signingUser,
// 			v0.ConditionNames_InSets.String():       signingUser,
// 			v0.ConditionNames_MultiSig.String():     trs.Witness,
// 		},
// 	}, policies)
//
// }
