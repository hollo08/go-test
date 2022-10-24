package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
)

func main(){
    files := []string{"ORDER_TRADING_20211130_95bde9bdedb1449480fadd4d3da28fb0.zip","ORDER_TRADING_20211123_1be5047d3a16483f8546c4ab9bb68634.zip","ORDER_TRADING_20211116_dc70dc8bd63241baab4388b0c5a96d6b.zip","ORDER_TRADING_20211109_7be01b3d401842a4a8cc33b0de59a52f.zip","ORDER_TRADING_20211102_d207a5a324e0403aa4d2adec79a12ac7.zip","ORDER_TRADING_20211026_bc582ab7d29245ac9aa01be872c3961d.zip","ORDER_TRADING_20211026_bfb7f17d77234a409d77e85833308597.zip","ORDER_TRADING_20211019_8c700081fd48406f99ddfbb0036a9d5c.zip","ORDER_TRADING_20211012_efeff496714d4249ab660ab78b8e0513.zip","ORDER_TRADING_20211005_eb15a1821c414d259f01b167b1effc7e.zip","ORDER_TRADING_20210928_11f0d6f8d63d48bca7670a57eb89dd71.zip","ORDER_TRADING_20210921_06bb93fc98d84cf2a25c1fcbe773db6f.zip","ORDER_TRADING_20210920_27a4e22e80184fd8bd424b746e2941e1.zip","ORDER_TRADING_20210914_3a2d62ead9074fc28e116febb2b05225.zip","ORDER_TRADING_20210913_b1ae4b6b97734b69b6aa289867067bde.zip","ORDER_TRADING_20210907_652ae163746e4ac0b95e1a7fc9b1d23a.zip","ORDER_TRADING_20210906_5f7327a72f004a79be9d9fd3e17ebbd2.zip","ORDER_TRADING_20210831_6b87bb2a5e264e5c90a2174f8446a13d.zip","ORDER_TRADING_20210830_452e81aae44c47fe8a3749796589ce25.zip","ORDER_TRADING_20210824_c59084cd366e4b5da99c86bd5571d623.zip","ORDER_TRADING_20210823_ec6db5ad576147e6954ccfd44803d7a3.zip","ORDER_TRADING_20210817_94fcfef7de9845229726fa282da49633.zip","ORDER_TRADING_20210816_0bd950e4d82c4bc99f007d221d3a3d10.zip","ORDER_TRADING_20210810_f2647f3cb9674bb7b5505236e8033f93.zip","ORDER_TRADING_20210809_6f87e9669c954c919223a3e79388f454.zip","ORDER_TRADING_20210803_a5d3bbf119b34069a7f690fb2620c0c7.zip","ORDER_TRADING_20210802_60a85f5dd32740ccb64cd3f98f07437f.zip","ORDER_TRADING_20210727_785920530bbe4ad0b96dca669ded399d.zip","ORDER_TRADING_20210726_866207195ba24456be46127139c9a7e2.zip","ORDER_TRADING_20210720_a09fe3ecec0946f886d00696a66265ec.zip","ORDER_TRADING_20210719_579ffe854e704c59839cde26bfc70c8f.zip","ORDER_TRADING_20210712_04b685541ec84ef2a884f66dce139aa9.zip","ORDER_TRADING_20210706_19baea4ea32d491c9b097f636bf09061.zip","ORDER_TRADING_20210705_568de633133b45408d9c30f264995396.zip","ORDER_TRADING_20210629_bbbca1e19463437faefa2e0c9b53fce9.zip","ORDER_TRADING_20210628_4a7b9c0372a742798a36267660913d85.zip","ORDER_TRADING_20210622_2d2958bae862434e8a9011241894efb7.zip","ORDER_TRADING_20210621_b501d2796ad44e9493de2de58aae1ce3.zip","ORDER_TRADING_20210615_c56ede2ae4be41b88bf03f23d8b9840f.zip","ORDER_TRADING_20210615_e10fc3ad8e9b41d79ac77784f5fa056e.zip","ORDER_TRADING_20210614_489049cd50224cf8b71a25ea6012597f.zip","ORDER_TRADING_20210608_8ce5522856674d45b4dcd1ad5ee4bb15.zip","ORDER_TRADING_20210608_427e651b7724407bbd8c6eac0ea1093a.zip","ORDER_TRADING_20210607_603b6bb4be9c4bd8ad07f1ba1b65a08c.zip","ORDER_TRADING_20210601_d2b6d4e3a4204baab2b30d5851d354a2.zip","ORDER_TRADING_20210601_73c62b72c66846dbbb0775519026e3b9.zip","ORDER_TRADING_20210531_ede89a02ab3349c6bb214bc8ca706982.zip","ORDER_TRADING_20210525_9602a7744e1e458a993cbd68f0dfadf8.zip","ORDER_TRADING_20210525_4177ad6fe2d245a8b6eeb1a5eeeea0e7.zip","ORDER_TRADING_20210524_070dc83f4c604aaea5416fe5ce06c43f.zip","ORDER_TRADING_20210518_55c21a45fd2d424dab40de560fdd43e5.zip","ORDER_TRADING_20210518_75e9f81f43c84c239850484595a2e7bf.zip","ORDER_TRADING_20210517_f4f7ba0f5bf54b86b80978d7690aa950.zip","ORDER_TRADING_20210511_3e4d9f70619c44a79e6da87cfd18804f.zip","ORDER_TRADING_20210511_cf519f10a7f248c3aa7644ca26ed58a7.zip","ORDER_TRADING_20210510_d0ec84c4affa47178d57d9394b9aaa6e.zip","ORDER_TRADING_20210504_5ed7abc02e9c41e4a9d3119dbf41556f.zip","ORDER_TRADING_20210503_0642831f18674a90858b9e5712affc44.zip","ORDER_TRADING_20210429_61b1a755741c48a89d9fc12248953900.zip","ORDER_TRADING_20210429_22416ac502fb482db8133d8c4c5eacf6.zip","ORDER_TRADING_20210429_36512fa438664b6a945007d2f69d73b3.zip","ORDER_TRADING_20210429_20b15fec12dc4c5e8b066da71327f3b5.zip","ORDER_TRADING_20210429_fbf5f29eea5642639300fc556e8b8f61.zip","ORDER_TRADING_20210429_06b64accd3b941c88977972defab5cb5.zip","ORDER_TRADING_20210429_25e2ebe556f443e6a21e15eccae7ebe6.zip","ORDER_TRADING_20210429_029f15bbe2e644b1a85fc8586b01c746.zip","ORDER_TRADING_20210429_62e04f5d155b4f8b92dfba948249c373.zip","ORDER_TRADING_20210429_c2a6f6d3a7aa4458be46cee55ca6a971.zip","ORDER_TRADING_20210429_68cffc2ca5f64edab83c3594cd851595.zip","ORDER_TRADING_20210429_661f55cef5f44a7e9b0c6395a99406b9.zip","ORDER_TRADING_20210429_31ea5c2fed714b5280471cb238b4c89f.zip","ORDER_TRADING_20210429_a3b3bd2a1a4f445992f2fae4a35ab0f9.zip","ORDER_TRADING_20210429_b547812f99704493a74b8d8c48fdbb60.zip","ORDER_TRADING_20210429_a3089d1d97c54b4c8993b9b808d3c6f5.zip","ORDER_TRADING_20210429_f6b7d172ae464be0a711ad36f6f427bc.zip","ORDER_TRADING_20210429_a8cd74d805e94bef896be90275fcf67c.zip","ORDER_TRADING_20210429_31d3e5f7f4464fdd8794c94ea7531235.zip","ORDER_TRADING_20210429_5579447535ad4e5493f21bac4b1fb53c.zip","ORDER_TRADING_20210429_38287ecd4f664cdcb3827b01b04f770b.zip","ORDER_TRADING_20210429_c2751f455fa14651be5bf6f8e4be036e.zip","ORDER_TRADING_20210429_71ccff1e36d742e29466a11eee0b7b40.zip","ORDER_TRADING_20210429_cfeca635cd404ae8b32b848b9bd4a2b2.zip","ORDER_TRADING_20210429_203bd3026936480699de1a261b02ebc5.zip","ORDER_TRADING_20210429_1ed66c9796b64aebaf4ba5018f687347.zip","ORDER_TRADING_20210429_e09265ccfc0c461dbe4c45a4d3c45b86.zip","ORDER_TRADING_20210429_2b9f94bfe4564b31b60ad4b1c0cd5ab8.zip","ORDER_TRADING_20210429_ae0ee393bdb5470db1ea32ed2671a5bb.zip","ORDER_TRADING_20210429_49b53f03f9fc4c30bafd34a7f2e3a6ae.zip","ORDER_TRADING_20210429_ecf527e0e2f348e8b0efdaa8630fa2b0.zip","ORDER_TRADING_20210429_c3490b5004c049c5b38c480b4219be58.zip","ORDER_TRADING_20210429_2ba9039c927d48b88609dc9fa65a6fd5.zip","ORDER_TRADING_20210429_f1c47ba7a2b54c04842a84d577d837a0.zip","ORDER_TRADING_20210429_f78ab970bd1d463db3ff130291e54478.zip","ORDER_TRADING_20210429_1893b27339f64b7f83dcd72077c4fc44.zip","ORDER_TRADING_20210429_ee78ed56c1814c978e512d12048bee35.zip","ORDER_TRADING_20210429_ee78f2a9a4e74d61bc7850c11108bc8f.zip","ORDER_TRADING_20210429_9b468cb5d4494ea68e32982823941d59.zip","ORDER_TRADING_20210429_4f64ef9bf45045499043da6dd5edaa3b.zip","ORDER_TRADING_20210429_726212b5c4ee4a6490afeb0ecca95b50.zip","ORDER_TRADING_20210429_f5d5138179804cb7b9da45fbbe1f5ad9.zip","ORDER_TRADING_20210429_0aca75226ba14816829e8f205949f5ba.zip","ORDER_TRADING_20210429_2369dbeeee104d1380164c16e2ec1a2a.zip","ORDER_TRADING_20210429_2f7a0032ed884d7383dd15f5d7bcde25.zip","ORDER_TRADING_20210429_d5fd790a259c487faed095bf6e7d5850.zip","ORDER_TRADING_20210429_a4df52eb60b544da8c2008b744277f1e.zip","ORDER_TRADING_20210429_025e0072787f43a2a8a4f4ecbd76b12a.zip","ORDER_TRADING_20210429_f76ca842610c45adafe791e9f2e3e2ae.zip","ORDER_TRADING_20210429_d59dcfcb59554457b1daf8033de0e5b8.zip","ORDER_TRADING_20210429_747d043820aa42b89f35d3457d87a645.zip","ORDER_TRADING_20210429_e8070f0d5adf43a2ac7b60bdee334990.zip","ORDER_TRADING_20210429_b5e9d7fdadc14d7d82294b6311b5cf3a.zip","ORDER_TRADING_20210429_914142af975a4fe69bcd53f32522709a.zip","ORDER_TRADING_20210429_b68dc8f1f23848eb901bcf945ec98946.zip","ORDER_TRADING_20210429_412fc3e1c7ba4c4bbb54325f4086bfba.zip","ORDER_TRADING_20210429_8ef33dacf5fb4f888313892b93b4f720.zip","ORDER_TRADING_20210429_e5dda6fa22844ef7a7cbc36a76cf0772.zip","ORDER_TRADING_20210429_3d923e19c28e4fe5b73f01d17cfa53d1.zip","ORDER_TRADING_20210429_1114e59536e5463eab34eb096ddbc60a.zip","ORDER_TRADING_20210429_f3505892bb684d9581e701c4744fb7ff.zip","ORDER_TRADING_20210429_c7cc8768caab422190b5f83ec38a57fb.zip","ORDER_TRADING_20210429_45e3c225f38a4f35a5946e9126f8aece.zip","ORDER_TRADING_20210429_12d5c72723e44c57a7f7abffaa8b1988.zip","ORDER_TRADING_20210429_297b238c889e4979a7fc1fe38d3237ac.zip","ORDER_TRADING_20210429_ea8792a3e5364b91a2c622be03acc4f6.zip","ORDER_TRADING_20210429_dfb5114aed4c41deb972900b820f04ef.zip","ORDER_TRADING_20210429_debcd7ba0cf14f908cb22acf74f2c029.zip","ORDER_TRADING_20210429_277f8469fcd64a24a94d6f7392f85d0b.zip","ORDER_TRADING_20210429_f3408c29c24644d2a92d0fef566b1539.zip","ORDER_TRADING_20210429_37b9ecb505fb4021bd838f8f4a042cfd.zip","ORDER_TRADING_20210429_0256ec01cc3b4a9b9c1e7ac6a2b96382.zip","ORDER_TRADING_20210429_12813d54207f435e8ff5f8a1c1b60c9f.zip","ORDER_TRADING_20210429_42cee63c86f64a00bd3d6700df311882.zip","ORDER_TRADING_20210429_08c5e4b117774ebbba07534d706204c6.zip","ORDER_TRADING_20210429_514eccab10844bdbac0201b2e9e3e436.zip","ORDER_TRADING_20200929_c055a11940af4001b9116849a9386ea7.zip","ORDER_TRADING_20200929_d7d57112d4dc4b049bbce5cde5ddaf54.zip","ORDER_TRADING_20200928_80958724826a46f4bfefd95afcfb7a52.zip","ORDER_TRADING_20200922_bfa6eb439db645fcb8583c4a57f74727.zip","ORDER_TRADING_20200922_9ee11832b636419c81d09b58d87fa20e.zip","ORDER_TRADING_20200921_2bf31b8215024fe991b2b3fa0af0defb.zip","ORDER_TRADING_20200915_4bf1997a43714f1a86f840878374ca6b.zip","ORDER_TRADING_20200915_66f02b7bb49c4b0db860e06e3f53e48c.zip","ORDER_TRADING_20200914_c5ddb8ea755340b78d0e7b0247667619.zip","ORDER_TRADING_20200824_750724626ad04614965d42a45085b9df.zip","ORDER_TRADING_20200824_1532bdbfa9d24c6192df5d3da76f63e7.zip","ORDER_TRADING_20200824_809f38ebe3c14bb2b01c8d6b16a104bd.zip","ORDER_TRADING_20200824_f939ecffdc5b4c5094e5592a80395713.zip","ORDER_TRADING_20200824_ad5ac3bb64404d368ffbf30c587c0607.zip","ORDER_TRADING_20200824_b4dae44f117f45509131a07fcdb824ed.zip","ORDER_TRADING_20200824_62b3be6583934747b1c069ba7bd0bcbc.zip","ORDER_TRADING_20200824_ebb93fea9b91468098ff52b0c6f36b1f.zip","ORDER_TRADING_20200824_e0ace87d39fa453faeec52ef4f00c1b5.zip","ORDER_TRADING_20200824_cdb7c0ddeb2d4430ad225da16f88ce02.zip","ORDER_TRADING_20200824_bf67130035684b8ab579ef5ff33e6e4d.zip","ORDER_TRADING_20200824_c5806001612945f49d9d8e181eb5f279.zip","ORDER_TRADING_20200824_80c35978de624299b3db776a0cc6bf0d.zip","ORDER_TRADING_20200824_8650dadadba14e2e8d384ccfdb3340b8.zip","ORDER_TRADING_20200824_1a777348ae864f988e14784bcda0dea0.zip","ORDER_TRADING_20200824_ee58405fb14545e195bcd65afc279bed.zip","ORDER_TRADING_20200824_5b1a35fd53564f32bac97249245f6cd1.zip","ORDER_TRADING_20200824_b820723f94954495ab7a77a1de2fbd6e.zip","ORDER_TRADING_20200824_6476fe304458422a96a47552f3ee4803.zip","ORDER_TRADING_20200824_7de9257b6b8540bb95af9562380b5b2c.zip"}
    imgPath := `F:\order\`
    for _, a := range files{
        fmt.Println("https://bolome-file.oss-cn-hangzhou.aliyuncs.com/"+a)
        imgUrl := "https://bolome-file.oss-cn-hangzhou.aliyuncs.com/"+a
        fileName := path.Base(imgUrl)
    
    
        res, err := http.Get(imgUrl)
        if err != nil {
            fmt.Println("A error occurred!")
            return
        }
        defer res.Body.Close()
        // 获得get请求响应的reader对象
        reader := bufio.NewReaderSize(res.Body, 32 * 1024)
    
    
        file, err := os.Create(imgPath + fileName)
        if err != nil {
            panic(err)
        }
        // 获得文件的writer对象
        writer := bufio.NewWriter(file)
    
        written, _ := io.Copy(writer, reader)
        fmt.Printf("Total length: %d", written)
    }
}