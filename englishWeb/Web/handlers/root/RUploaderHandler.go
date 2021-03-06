package root

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
	"io"
	"os"
	"strings"
	"time"
)

const (
	outtimes = "Error:"
)

type RUploaderHandler struct {
	base.BaseHandler
}

func (self *RUploaderHandler) Get() {
	if sess_role, _ := self.GetSession("userrole").(int64); sess_role != -1000 {
		self.Ctx.WriteString(outtimes + "请你重新登录，现已超时操作！")
	} else {
		self.TplNames = "root/uploadify/index.html"
		self.Render()
	}

}

func (self *RUploaderHandler) Post() {
	/*
		//TODO: Validate the file type

	*/

	if sess_role, _ := self.GetSession("userrole").(int64); sess_role != -1000 {
		_, handler, _ := self.GetFile("uploadfile")

		if handler != nil {
			self.Ctx.WriteString(outtimes + "上传“ " + handler.Filename + " ”失败，请你重新登录，现已超时操作！")
		} else {
			self.Ctx.WriteString(outtimes + "请你重新登录，现已超时操作！")
		}
	} else {
		targetFolder := "/archives/upload/"

		file, handler, e := self.GetFile("uploadfile")
		pid, _ := self.GetInt("pid")

		if e != nil {
			self.Data["MsgErr"] = "0"
		} else {

			if handler != nil {
				ext := "." + strings.Split(handler.Filename, ".")[1]
				filename := common.MD5(time.Now().String()) + ext

				path := targetFolder + time.Now().Format("2006/01/02/")

				os.MkdirAll("."+path, 0644)
				path = path + filename
				f, err := os.OpenFile("."+path, os.O_WRONLY|os.O_CREATE, 0644)
				defer f.Close()

				if err != nil {
					self.Data["MsgErr"] = "0"
				} else {
					io.Copy(f, file)
					defer file.Close()
					input_file := "." + path
					output_file := "." + path
					output_size := "534"
					output_align := "center"
					background := "black"
					common.Thumbnail(input_file, output_file, output_size, output_align, background)
					hash := common.Filehash(output_file)
					fileInfo, err := os.Stat(output_file)
					var fsize int64 = 0
					if err == nil {
						fsize = fileInfo.Size() / 1024
					}

					self.Data["MsgErr"] = "<img src=\"" + path + "\" alt=\"" + hash + "\" />"

					bll.SetFile(0, pid, 0, handler.Filename, "", hash, path, "", fsize)
				}

			} else {
				self.Data["MsgErr"] = "0"
			}
		}

		self.Ctx.WriteString(self.Data["MsgErr"].(string))
	}
}
