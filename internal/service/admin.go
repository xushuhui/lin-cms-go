package service

// func GetUsers(c *fiber.Ctx) error {
// 	var req api.ListUserRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}

// 	data, err := biz.GetUsers(c.Context(), int(req.GroupId), core.GetPage(c), core.GetSize(c))
// 	if err != nil {
// 		return err
// 	}
// 	return core.SetData(c, data)
// }

// func ChangeUserPassword(c *fiber.Ctx) error {
// 	var req api.UpdateUserPasswordRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}

// 	err := biz.ChangeUserPassword(c.Context(), int(req.Id), req.NewPassword)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func DeleteUser(c *fiber.Ctx) error {
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}
// 	err = biz.DeleteUser(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func UpdateUser(c *fiber.Ctx) error {
// 	var req api.UpdateUserRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}

// 	err := biz.UpdateUser(c.Context(), int(req.Id), req.GroupIds)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }
