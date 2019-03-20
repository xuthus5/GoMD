package tools

import (
	"io"
	"io/ioutil"
	"os"
)

/*************
自定义文件操作
**************/

/* 文件删除 */
func FileRemove(path string) error {
	return os.Remove(path)
}

/* 检测目录是否存在 */
func CheckFilePath(path string) error {
	_, err := os.Stat(path) //检测路径状态
	if err == nil {
		return nil //没有错误 表明文件夹路径存在
	}
	return err
}

/* 创建目录 */
func DirCreate(path string) error {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err //创建目录失败
	}
	return nil
}

/* 检测目录是否存在 并且创建目录 */
func CheckAndDirCreate(path string) error {
	_, err := os.Stat(path) //检测路径状态
	if err == nil {
		return nil //没有错误 表明文件夹路径存在 返回true nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

/* 文件移动 */
func FileMove(source, target string) error {
	srcFile, err := os.Open(source) //打开源文件
	if err != nil {
		return err
	}
	defer srcFile.Close()
	tagFile, err := os.Create(target) //打开目标文件
	if err != nil {
		return err
	}
	defer tagFile.Close()
	_, err = io.Copy(tagFile, srcFile) //文件拷贝
	if err != nil {
		return err //操作失败
	}
	code := FileRemove(source) //源文件删除
	if code != nil {
		return code //操作失败
	} else {
		return nil
	}
}

/* 遍历目录获取文件夹 */
func ErgodicPathGetDir(path string, theme bool) []string {
	//特殊要求 遍历视图目录 获取主题
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}
	}
	dir := []string{}
	for _, fi := range rd {
		if fi.IsDir() && theme == true && fi.Name() != "admin" {
			dir = append(dir, fi.Name())
		}
		if fi.IsDir() && theme == false {
			dir = append(dir, fi.Name())
		}
	}
	return dir
}

/* 写入文件 */
func WriteFile(filePath, fileContent string) error {
	var f *os.File
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, fileContent)
	return err
}
