package service

// func GetBooks(c *fiber.Ctx) error {
// 	// TODO book 接口少了权限判断

// 	size := core.GetSize(c)
// 	page := core.GetPage(c)
// 	data, err := biz.GetBookAll(c.Context(), page, size)
// 	if err != nil {
// 		return err
// 	}
// 	total, err := biz.GetBookTotal(c.Context())
// 	if err != nil {
// 		return err
// 	}
// 	core.SetPage(c, data, total)
// 	return nil
// }

// func UpdateBook(c *fiber.Ctx) error {
// 	var req api.UpdateBookRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}

// 	err = biz.UpdateBook(c.Context(), id, &req)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func CreateBook(c *fiber.Ctx) error {
// 	var req api.CreateBookRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}
// 	err := biz.CreateBook(c.Context(), &req)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func DeleteBook(c *fiber.Ctx) error {
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}
// 	err = biz.DeleteBook(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func GetBook(c *fiber.Ctx) error {
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}
// 	data, err := biz.GetBook(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}
// 	core.SetData(c, data)
// 	return nil
// }
