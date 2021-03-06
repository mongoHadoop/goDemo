﻿package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main6738() {
	fmt.Println("---------------------tar写入缓存------------------------")
	buf, e_tar := os.OpenFile("main/aa.txt", os.O_RDWR, 777)
	check_EncDec_err(e_tar)

	//不一要写入文件，也可以放到缓存区
	//创建一个缓冲区以将存档写入其中。
	//buf := new(bytes.Buffer)
	//创建一个新的tar存档。
	// NewWriter创建一个写入w的新Writer。
	tw := tar.NewWriter(buf)
	//将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for _, file := range files {
		//header代表tar归档文件中的单个标头。
		//某些字段可能未填充。
		//
		//为了向前兼容，用户从Reader中检索Header，然后以某种方式对其进行突变，然后将其传递回Writer.WriteHeader应该通过创建新Header并复制他们想要保留的字段来实现。
		//type Header struct {
		//	// Typeflag是标题条目的类型。
		//	//根据Name中是否存在尾部斜杠，零值会自动提升为TypeReg或TypeDir。
		//	Typeflag byte
		//
		//	Name     string //文件入口名称
		//	Linkname string //链接的目标名称（对TypeLink或TypeSymlink有效）
		//
		//	Size  int64  //逻辑文件大小（以字节为单位）
		//	Mode  int64  //权限和模式位
		//	Uid   int    //所有者的用户标识
		//	Gid   int    //所有者的组ID
		//	Uname string //所有者的用户名
		//	Gname string //所有者的组名
		//
		// //如果未指定Format，则Writer.WriteHeader将ModTime舍入到最接近的秒数，并忽略AccessTime和ChangeTime字段。
		// //要使用AccessTime或ChangeTime，请将Format指定为PAX或GNU。
		// //要使用亚秒级分辨率，请将格式指定为PAX。
		//	ModTime    time.Time //修改时间
		//	AccessTime time.Time //访问时间（需要PAX或GNU支持）
		//	ChangeTime time.Time //更改时间（需要PAX或GNU支持）
		//
		//	Devmajor int64 //主设备号（对TypeChar或TypeBlock有效）
		//	Devminor int64 //次设备号（对TypeChar或TypeBlock有效）
		//
		//	//Xattrs将扩展属性作为PAX记录存储在“ SCHILY.xattr”下的命名空间。
		//	//以下在语义上是等效的：
		//	//h.Xattrs [key] = value
		//	//h.PAXRecords [“ SCHILY.xattr。” + key] =值
		//
		//	//调用Writer.WriteHeader时，Xattrs的内容将优先于PAXRecords中的内容。
		//	//不推荐使用：请改用PAXRecords。
		//	Xattrs map[string]string
		//
		// // PAXRecords是PAX扩展头记录的映射。
		// //用户定义的记录应具有以下格式的键：
		// // VENDOR.keyword其中VENDOR是所有大写形式的名称空间，而关键字可能不包含'='字符（例如，“ GOLANG.pkg.version”）。
		// //键和值应为非空的UTF-8字符串。
		// //
		// //调用Writer.WriteHeader时，从Header中其他字段派生的PAX记录优先于PAXRecords。
		//	PAXRecords map[string]string
		//
		// // Format指定tar标头的格式。
		// //这是由Reader.Next设置的，是对格式的最大努力。
		// //由于Reader可以自由读取一些不兼容的文件，因此可能是FormatUnknown。
		// //
		// //如果在调用Writer.WriteHeader时未指定格式，则它将使用能够对该Header进行编码的第一种格式（按USTAR，PAX，GNU的顺序）（请参见Format）。
		//	Format Format
		//}

		hdr := &tar.Header{
			Name: file.Name,	//文件入口名称
			Size: int64(len(file.Body)),	//逻辑文件大小（以字节为单位）
		}
		// WriteHeader写入hdr并准备接受文件的内容。
		// Header.Size确定可以为下一个文件写入多少字节。
		//如果当前文件未完全写入，则返回错误。
		//这会在写入标题之前隐式刷新所有必要的填充。
		if err := tw.WriteHeader(hdr); err != nil {//首先写入header,相当于创建一个新的空的文件元数据（文件可以由多个元数据组成），并且准备接受文件的内容,
			log.Fatalln(err)
		}
		//Write（）写入tar存档中的当前文件。
		//如果在WriteHeader之后写入了超过Header.Size字节，则Write返回错误ErrWriteTooLong。
		//
		//在特殊类型（如TypeLink，TypeSymlink，TypeChar，TypeBlock，TypeDir和TypeFifo）上调用Write返回（0，ErrWriteTooLong），无论Header.Size声明什么。
		if _, err := tw.Write([]byte(file.Body)); err != nil {//开始往header创建的新的空文件上面写入文件的内容
			log.Fatalln(err)
		}
	}
	//确保在关闭时检查错误。
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	//写入buf
	//fmt.Println("tar写入缓存中的字节数据为：",buf.Bytes())
	//fmt.Printf("tar写入缓存中的字符串数据为：\n%v\n",buf.String())
	//写入文件
	buf, e_tar1111 := os.OpenFile("main/aa.txt", os.O_RDWR, 777)//这里的buf更改了上面的额buf的指向，上面的buf已经关闭了，无法再次读取，所以这里需要重新读取上面的buf的文件的内容
	check_EncDec_err(e_tar1111)
	info, _ := buf.Stat()
	byte:=make([]byte,info.Size())//通过文件大小来创建相应大小的byte
	n, e_read := buf.Read(byte)
	if e_read != nil {
		log.Fatalln("1111111111",e_read)
	}
	fmt.Println(n)
	fmt.Println("tar写入缓存中的字节数据为：",byte)
	fmt.Printf("tar写入缓存中的字符串数据为：\n%v\n",string(byte))


	fmt.Println("---------------------tar读取缓存------------------------")
	//打开tar存档以供读取。
	//r := bytes.NewReader(buf.Bytes())//将上面写入的内容读取出来,如果是写入文件袋额话，就不用这样获取

	buf, e_tar222 := os.OpenFile("main/aa.txt", os.O_RDWR, 777)//这里的buf再次更改了上面的额buf的指向，上面的buf已经关闭了，无法再次读取，所以这里需要重新读取上面的buf的文件的内容
	check_EncDec_err(e_tar222)
	// Reader提供对tar存档内容的顺序访问。
	// Reader.Next前进到存档中的下一个文件（包括第一个文件），然后Reader可被视为io.Reader来访问文件的数据。
	//type Reader struct {
	//	r    io.Reader
	//	pad  int64      //当前文件输入后的填充量（忽略）
	//	curr fileReader //当前文件输入的阅读器,这个接口由2个接口和一个方法组成：io.Reader，fileState（fileState跟踪当前文件剩余的逻辑（包括稀疏漏洞）和物理（实际在tar存档中）字节的数量。），WriteTo(io.Writer) (int64, error)
	//	blk  block      //用作临时本地存储的缓冲区，底层是容量为512的数组[512]byte
	//
	//	// err是一个持久错误。
	//	//确保Reader错误是粘性的，这是Reader的每个导出方法的责任。
	//	err error
	//}
	tr := tar.NewReader(buf)// NewReader从r创建一个新的Reader读取器。
	//遍历档案中的文件。
	for {
		//Next()前进到tar存档中的下一个条目。
		// Header.Size确定下一个文件可以读取多少个字节。
		//当前文件中的所有剩余数据将被自动丢弃。
		// io.EOF在输入末尾返回。
		//在外部，Next遍历tar存档，就好像它是一系列文件一样。 在内部，tar格式通常使用伪造的“文件”来添加描述下一个文件的元数据。
		// 这些元数据“文件”通常不应在外部可见。 这样，此循环遍历一个或多个“头文件”，直到找到“正常文件”为止。
		hdr, err := tr.Next()
		if err == io.EOF {
			// tar存档的结尾,break结束当前的循环
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("----下面是通过header对象获取os.FileInfo对象的一些信息----")
		fmt.Println("--------FileInfo对象:",hdr.FileInfo())
		fmt.Println("--------FileInfo对象.Name():",hdr.FileInfo().Name())
		fmt.Println("--------FileInfo对象.Size():",hdr.FileInfo().Size())
		fmt.Println("--------FileInfo对象.Mode():",hdr.FileInfo().Mode())
		fmt.Println("--------FileInfo对象.ModTime():",hdr.FileInfo().ModTime())
		fmt.Println("--------FileInfo对象.IsDir():",hdr.FileInfo().IsDir())
		fmt.Println("--------FileInfo对象.Sys():",hdr.FileInfo().Sys())
		fmt.Println()

		fmt.Println("----下面是获取当前缓存文件的header对象的一些信息----")
		fmt.Println("Typeflag是标题条目的类型：",hdr.Typeflag)// Typeflag是标题条目的类型。根据Name中是否存在尾部斜杠，零值会自动提升为TypeReg或TypeDir。
		fmt.Println("文件入口名称：",hdr.Name)//文件入口名称
		fmt.Println("链接的目标名称：",hdr.Linkname)//链接的目标名称（对TypeLink或TypeSymlink有效）
		fmt.Println("逻辑文件大小：",hdr.Size)//逻辑文件大小（以字节为单位）
		fmt.Println("权限和模式位：",hdr.Mode)//权限和模式位
		fmt.Println("所有者的用户标识：",hdr.Uid)//所有者的用户标识
		fmt.Println("所有者的组ID：",hdr.Gid)//所有者的组ID
		fmt.Println("所有者的用户名：",hdr.Uname)//所有者的用户名
		fmt.Println("所有者的组名：",hdr.Gname)//所有者的组名
		fmt.Println("修改时间：",hdr.ModTime)//修改时间
		fmt.Println("访问时间：",hdr.AccessTime)//访问时间（需要PAX或GNU支持）
		fmt.Println("更改时间：",hdr.ChangeTime)//更改时间（需要PAX或GNU支持）
		fmt.Println("主设备号：",hdr.Devmajor)//主设备号（对TypeChar或TypeBlock有效）
		fmt.Println("次设备号：",hdr.Devminor)//次设备号（对TypeChar或TypeBlock有效）
		fmt.Println("PAX扩展头记录的映射：",hdr.PAXRecords)// PAXRecords是PAX扩展头记录的映射。
		fmt.Println("tar标头的格式：",hdr.Format)// Format指定tar标头的格式。
		fmt.Println("-----------------------------")

		fmt.Printf("Contents of %s:\n", hdr.Name)
		//将副本从src复制到dst，直到在src上达到EOF或发生错误。 它返回复制的字节数和复制时遇到的第一个错误（如果有）。
		//成功的Copy返回err == nil，而不是err == EOF。
		//因为复制被定义为从src读取直到EOF，所以它不会将读取的EOF视为要报告的错误。
		//如果src实现WriterTo接口，则通过调用src.WriteTo（dst）实现该副本。
		//否则，如果dst实现了ReaderFrom接口，则通过调用dst.ReadFrom（src）实现该副本。
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()

		}

	//输出：
	//	---------------------tar写入缓存------------------------
	//	tar写入缓存中的字节数据为： [114 101 97 100 109 101 46 116 120 116 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 48 48 52 54
	//	0 48 48 48 48 48 48 48 48 48 48 48 0 48 49 49 49 54 53 0 32 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 117 115 116 97 114 0 48 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 84 104 105 115 32 97 114 99 104 105
	//	118 101 32 99 111 110 116 97 105 110 115 32 115 111 109 101 32 116 101 120 116 32 102 105 108 101 115 46 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 103 111 112 104 101 114 46 116 120 116 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 48
	//	48 52 51 0 48 48 48 48 48 48 48 48 48 48 48 0 48 49 49 50 49 49 0 32 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 117 115 116 97 114 0 48 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 71 111 112 104 101 114 32 110 97 109 101 115 58 10 71 101 111 114 103 101 10
	//	71 101 111 102 102 114 101 121 10 71 111 110 122 111 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 116 111 100 111 46 116 120 116 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 48 48 48 48 48 48 0
	//	48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 48 48 51 52 0 48 48 48 48 48 48 48 48 48 48 48 0 48 49
	//	48 54 55 50 0 32 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 117 115 116 97 114 0 48 48
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 48 48 48 48 48 48 48 0 48 48 48 48 48 48 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 71 101 116 32 97 110 105 109 97 108 32 104 97 110 100 108 105 110 103 32 108 105 99 101 110 99 101 46 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	tar写入缓存中的字符串数据为：
	//	readme.txt                                                                                          0000000 0000000 0000000 00000000046 00000000000 011165  0                                                                                                    ustar 00                                                                0000000 0000000                                                                                                                                                                        This archive contains some text files.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          gopher.txt                                                                                          0000000 0000000 0000000 00000000043 00000000000 011211  0                                                                                                    ustar 00                                                                0000000 0000000                                                                                                                                                                        Gopher names:
	//	George
	//	Geoffrey
	//	Gonzo                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             todo.txt                                                                                            0000000 0000000 0000000 00000000034 00000000000 010672  0                                                                                                    ustar 00                                                                0000000 0000000                                                                                                                                                                        Get animal handling licence.
	//---------------------tar读取缓存------------------------
	//----下面是通过header对象获取os.FileInfo对象的一些信息----
	//--------FileInfo对象: {0xc00009c1c0}
	//--------FileInfo对象.Name(): readme.txt
	//--------FileInfo对象.Size(): 38
	//--------FileInfo对象.Mode(): ----------
	//--------FileInfo对象.ModTime(): 1970-01-01 08:00:00 +0800 CST
	//--------FileInfo对象.IsDir(): false
	//--------FileInfo对象.Sys(): &{48 readme.txt  38 0 0 0   1970-01-01 08:00:00 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] USTAR}
	//
	//----下面是获取当前缓存文件的header对象的一些信息----
	//Typeflag是标题条目的类型： 48
	//文件入口名称： readme.txt
	//链接的目标名称：
	//逻辑文件大小： 38
	//权限和模式位： 0
	//所有者的用户标识： 0
	//所有者的组ID： 0
	//所有者的用户名：
	//所有者的组名：
	//修改时间： 1970-01-01 08:00:00 +0800 CST
	//访问时间： 0001-01-01 00:00:00 +0000 UTC
	//更改时间： 0001-01-01 00:00:00 +0000 UTC
	//主设备号： 0
	//次设备号： 0
	//PAX扩展头记录的映射： map[]
	//tar标头的格式： USTAR
	//-----------------------------
	//	Contents of readme.txt:
	//This archive contains some text files.
	//----下面是通过header对象获取os.FileInfo对象的一些信息----
	//--------FileInfo对象: {0xc0001420e0}
	//--------FileInfo对象.Name(): gopher.txt
	//--------FileInfo对象.Size(): 35
	//--------FileInfo对象.Mode(): ----------
	//--------FileInfo对象.ModTime(): 1970-01-01 08:00:00 +0800 CST
	//--------FileInfo对象.IsDir(): false
	//--------FileInfo对象.Sys(): &{48 gopher.txt  35 0 0 0   1970-01-01 08:00:00 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] USTAR}
	//
	//----下面是获取当前缓存文件的header对象的一些信息----
	//Typeflag是标题条目的类型： 48
	//文件入口名称： gopher.txt
	//链接的目标名称：
	//逻辑文件大小： 35
	//权限和模式位： 0
	//所有者的用户标识： 0
	//所有者的组ID： 0
	//所有者的用户名：
	//所有者的组名：
	//修改时间： 1970-01-01 08:00:00 +0800 CST
	//访问时间： 0001-01-01 00:00:00 +0000 UTC
	//更改时间： 0001-01-01 00:00:00 +0000 UTC
	//主设备号： 0
	//次设备号： 0
	//PAX扩展头记录的映射： map[]
	//tar标头的格式： USTAR
	//-----------------------------
	//	Contents of gopher.txt:
	//Gopher names:
	//George
	//Geoffrey
	//Gonzo
	//----下面是通过header对象获取os.FileInfo对象的一些信息----
	//--------FileInfo对象: {0xc0001422a0}
	//--------FileInfo对象.Name(): todo.txt
	//--------FileInfo对象.Size(): 28
	//--------FileInfo对象.Mode(): ----------
	//--------FileInfo对象.ModTime(): 1970-01-01 08:00:00 +0800 CST
	//--------FileInfo对象.IsDir(): false
	//--------FileInfo对象.Sys(): &{48 todo.txt  28 0 0 0   1970-01-01 08:00:00 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] USTAR}
	//
	//----下面是获取当前缓存文件的header对象的一些信息----
	//Typeflag是标题条目的类型： 48
	//文件入口名称： todo.txt
	//链接的目标名称：
	//逻辑文件大小： 28
	//权限和模式位： 0
	//所有者的用户标识： 0
	//所有者的组ID： 0
	//所有者的用户名：
	//所有者的组名：
	//修改时间： 1970-01-01 08:00:00 +0800 CST
	//访问时间： 0001-01-01 00:00:00 +0000 UTC
	//更改时间： 0001-01-01 00:00:00 +0000 UTC
	//主设备号： 0
	//次设备号： 0
	//PAX扩展头记录的映射： map[]
	//tar标头的格式： USTAR
	//-----------------------------
	//	Contents of todo.txt:
	//Get animal handling licence.
	//---------------------------------------
	//虽然我意图想通过写入文件来显示文件的信息内容，比如显示不同的权限和模式位，所有者的组ID等等

	fmt.Println("---------------------------------------")
}

































