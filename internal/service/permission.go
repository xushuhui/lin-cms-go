package service

// func GetAllPermissions(c *fiber.Ctx) error {

// 	data, err := biz.GetAllPermissions(c.Context())
// 	if err != nil {

// 		return err
// 	}
// 	return core.SetData(c, data)

// }

// //TODO: 没找到必须实现该接口的业务场景，而且功能和分配多个权限重复，开发待定
// func DispatchPermission(c *fiber.Ctx) error {

// 	return nil
// }

// func DispatchPermissions(c *fiber.Ctx) error {
// 	var req api.DispatchPermissionsRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}
// 	err := biz.DispatchPermissions(c.Context(), req.GroupId, req.PermissionIds)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)

// }

// func RemovePermissions(c *fiber.Ctx) error {
// 	var req request.RemovePermissions
// 	if err := core.ParseRequest(c, &req); err != nil {

// 		return err
// 	}
// 	err := biz.RemovePermissions(c.Context(), req.GroupId, req.PermissionIds)
// 	if err != nil {

// 		return err
// 	}
// 	return core.SuccessResp(c)

// }
