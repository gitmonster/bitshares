//this file is generated by gen [DO NOT EDIT]
//operation sample data for OperationTypeBlindTransfer

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeBlindTransfer] = sampleDataBlindTransferOperation
}

var sampleDataBlindTransferOperation = `{
  "fee": {
    "amount": 3000000,
    "asset_id": "1.3.0"
  },
  "inputs": [
    {
      "commitment": "0293ad9a2a14174a4fcd22b58fc662e878a03407727ca6afdb31d80aea414e185f",
      "owner": {
        "account_auths": [],
        "address_auths": [],
        "key_auths": [
          [
            "BTS6Hx8UT2pbejyZcj5mNoCKuHpaezuFNcwyGUVic1U4fX75NBRGN",
            1
          ]
        ],
        "weight_threshold": 1
      }
    }
  ],
  "outputs": [
    {
      "commitment": "0261d62e65662bfb34d0d9a61731ec1993be08014e245e8d1ef40b4258b115d952",
      "owner": {
        "account_auths": [],
        "address_auths": [],
        "key_auths": [
          [
            "BTS8C78RYB6zJCoQG5KbFVB7MbYmgZXGVTDHkhLQrfEdKgYGRnQni",
            1
          ]
        ],
        "weight_threshold": 1
      },
      "range_proof": "4016cb035317bc02ccd291c928ab80fd44d0163fd3afa351405b96bedc73a5a753a58323496f4df914ebea0d13db1f9ad8025820ea838c93c260645db82f828948b5515f0a678c387bab979c7c3713b4a754068479c270b21912fc3bf0702badf4392321495a0578d27d44f219d5b1ca56582298c8a3b3ddd334086ef07d222354e11ba40f59b5a64202bc1e6bf1855810192355c33973b6ef68462f811dcbc39eed713003bb90e322b05f2663a99417a6db15cc1bd5f1a38576d83441d022c804dd33be04363ba53e818b0f6efc572e8fc328e60c592280d8d448fb218e43138bead15018d82b93f8c9435f99ef92e90efa90fc704fe92c1207bac2ec39596058d7fd65daa39fbf610abe7828397a88e879fd9625c9312fd5e235ed9f93094e5f700aec723878aefcb2f4690b487584878668288f8161f281ace0aa3deb83b0e7bb305460029d8137d2404d2e9cefca669b77ca10cf0b87f6b5fcd290f2b6dbcd6608d6eec55e479978126bef8ebc61bc31c69e1a7106dade5614831f3035a6df6bfe015f0cec693ff6c89bb86cc727633a6c1ae7e0364a10b7aad10f3eea4424a308055172e285c476d8432fa4b090ad7f954a851efe0184950636464ac5e8fbf4e515747745ff3e338b3250159f55e64851a4f5fefde9a6719d96c6e7f55351042ec084b880f2d1b94b2e2bd934a3aa362bbe169f1cf969fe13c4b5447c17b3cea717beb51baa5b0ced6c3d2040186e8adbdb31cc94e58f846f4bbfa4756acf757ab7dc048171c51c28b8fb6142d91a4100b6d0fd1b79d918ffaaf45a317d0bb126d483b3fb214546ac3e03de1e5ea8a754c6fcda0e6ce5a1ee9acc72d920246dafbc7c0d2c37499b9f30fb6bb4c1944f2c1dbb75d9fa580f9a408cedbd813a4c7e15bb280b0abb84efb0fef7c4a09377cb4d0eafba3bb54a4ad09872983fe86d44bf2c4e7ef58caa3680f0bfdffe0ae7654754c8078ee4434d093662eaf41b2ee190a15625935515df827a002d165d56710e333a306b66721996b5d8f1d399f8024a45fca60d747364f3d8bc050e893d2e082db6573e1eae009fb9c946201253db71d4228af35193b2e0959cc3dba2857c1471549e8344905f00e1bfcd2427fb5f884e8ff5310a89d902444f1c4870e473b55efae3d89c4eaad110e70f3b7ab7c471adffe5ace5f271a0458f82641b7be49c01399f2cd9b94247ba748e5b98b96df9f725adacb382c81143e79cd4709a9ab078d72c54f8bcf63d934bb2230def5e56476e5e0cf50535377c7aa90b5f12bf85fbcdfef02ecd410a07faaf4c93de10f025f5bafa54cf05653c661a3b8feec83b168a30e7dc0908ee1ca9936228338610a686f5680163d38c7834667de35d9a05d5c31c13e6d6aad81544367cd7b8a2a5a594a1321fd7c7155bd8d7a1baf7fe6305886af628f1615e51bfc20b91b448b318a08f67027fb8df4b0f176da82b476dba1705b3c3d8f8d53fd01b6da0ec99915190bad52d6b112a76c478c9562acfcee01b39ad32fa6919adc95a9929390518253565fb065be512dc2114258c1294eba14b6b8073dd23557f9e4af5f4edd023c217235b38a9bf44ab7cba9602a417dae3b5cd518a0821aedb54ed2ab1889a032464bc45cd4581e2f4c3c630e566abb727583c01173dae01e25c880e96608eb7f9fae69f86283c3dca7b33603e697ae301cf14722b85dad6d32ac2d77b5d8532e1f4e011ca5786718e77ebc2f768a68f36d851b17d7e24523cd8f4faf8cb75e5f8ca94ee1fcc731274023d5af192bfc4c758152ab61360fd1f91621d34524093e14c99a8d1c963da2a6ee39808d665feb6bf741d9b379cf70b1ee03a02c510bf45023a7137de5d20e1c592cbb83fc97b0041d0c99533756710e35cd7a3b513df001e661b454c57c0ed2e42e6504e69916e9e51ffaabfcc463799a7b0b4027230883c49742c53c4daa1016fa6e410260caa1865909e800e71dcd3a58ab37cf680f9675d174cf1fcc664f1ae285346abed972b73032addb4f83ca40c896aa31283cd3c1d3e18cf00f0d800b7dff14e3c8dcb0025a95c787f3c101f2634bf00ffb3ee454073e36646557adc4a94cd609ba4f551116d3e104a8567ebde56fe27b87948cf176dd9a453435d5fca75abf3cd208bc9bc78a19bce82b68c69b5bfe63dd49569873228fff3acae8fb4e5d9f7be9e0955b4cc7d6c3b05d15af53623022c37c6210b2ec16b3a8b3fa4a8b60ba0e9c78be0764f74bfa68d7778e9cbcbb50b1c200ae218382a0556f74856be997192d01b55d0994179e1725d7eed3158c27dc5d6f7d46f5e4f5847fef19cc83e426dd35c0c4612f44a9e0098daef6d40adf66a500c891babd75f628ba1c7d7f2741334bcd887c9ae3e4239a1eb917b2fd6a692ddd674d5bf7544a2b5e51d6578200ab55b9dbb284ef7ceab5317db26cf7522e8052cfeedcef9df39ee1b7809e1f08fae3d10fba07506a55194c77db898727949ed6b63303b70b8163e2c15f0fa10b5cdd1ac976e5272c7c84a2fbfb842d2437144b6ebdff2d97b2833487a73dcb33c5bb2c66ffa920392916ba5ff3fc568170646a8f43c6b095cda43f320c4536878aad4ce0211137bd349e8030657e8302d8e"
    },
    {
      "commitment": "02fee663b188f809d8f476eb17757ac7e8f9bf28c694f9007dde1b90d6d3e72afb",
      "owner": {
        "account_auths": [],
        "address_auths": [],
        "key_auths": [],
        "weight_threshold": 0
      },
      "range_proof": "4015f902fb98da6d0de41a851d73eacc61f467020953f68167a1129589ca23f86cc012edde938a1df852d7d5bee1143a82fb9476a9fbd99e3b45fe419b94b8af7e88aefe003caa7335a745cdc135048d675f9e3ea58ddfb8c8442c1226804446d356198c90d8de109fc273c311307465a63dae7abc35910847ba0111b810c653488fcce576d524b7a907f7d3f9464a4ad309cb67fe9a4e6210c3905e131e826b47247c05e4214d702d2d50e49969042eda5dc76c51763858c483e92efc4c2a2cd789a884cbef7417d876280ef53360410526f056d033220d5cdbefecc7db024896221827807f85ae9dd87d0b80ddceaee95f38aa17324e8e1960004291db3e9bac7af24b459952584871df0901b06fa7b3f902abdc7bae8dd898c8b9e38753e4b6f2826ad7bd2a55525beb127cfddd91385c3b9a269d54b56117624e24c48fc9029085852a91387a2ea59ad44600119d39be7abaf2548099c2667ea62fdc70de52c4ef61285eca4eee49b81e1d366630850e36b65e2183bd55bc7af7bad769399b22e112981a33a6426a14899018aa3f112f00c5776d254aa4ea363e3f58f6f839fc66cf3db438b4651b599aae410529e49a83dfbada7cb07246d467b382cfa6fc774871d22a2cd606b8dc11a0f217e5469ca46c4d76e0e2d9064d09843e2cb96ee72d37a910da44886afb952c3a7d1c9fb07c4cec657b35a012818eafb5ecf23ce4a6d577a7170ccaf11755e194e1f9fc1b5e0b3630cbca9914c326cabb8911c3e69fad5ec9b82c3cc8cd92d822e206f4e26dcd702a40efef2512924ec252eb308b036470da41b6378f8c07ed5d3e2d9d746cd4ff9270af01af7884ef8fa11f70465d375d01256a3f2d9959a006a10d9681a3faa1d56c63470214292c88668f76f0f863da47f1a7d658b3b0cf6aee416a5f9eb7e149397c11ec28c835b76ea3fe3bdde58479385099098a37b283e704f572e091c9ea0fd32a67e59d2fb34df95354bad8cd97e7d35f1e94583aa0728f36194db4f1ebb1c9c956a1ac6fa46ac01a141c4655e7e49d741b2d0220ae9e996ae031970745d020eff86a31604cacfc9c2086a813e25ce2c3e6f087347ccbbfc675ff448140dba866ff4b3e63115550d981ecf431e1857f3d5dc31f2b3be7d48dde59064402995894edfdd8518953237eef03277fe4f7a5ae4e6008f2eba531fbaad023494b1574dd251f436178fe1dea50682c3e6cd52cede5b9d0f4e407ae4dd840db64fc7b6e256b8d9a13f115f91600e961848a0a22340fbd08ff622d911adc480f62f5e515531dd587aa816b30adec0a9efc9617a5830b88f449e684b972c8c4cbfc05fb34b7b72dad841a891e6f2619fe1e22e864babfdcde9c3cf83f5a6d1f2dcbdcf3a5f1b5e594bf7da8fa975a56effd65f6dc6321a03edfd31621534515d3d256c8df1d33179fbdc36aeb64f7e47c7bdb7d856dfcf3530f12776d0bc08f150453603e8647cecfee92e96e928a6b1f817f94faacdba836009d2b61ba6da23876e4791b227fe3d2d8954f981ae363c2793e68c6fc709f0cd61030897cf612e97c63cd89cf7fde647db8cb19a95a51e137fa12800e646d19811019d8175be2f411ea8e4459e530311cbe42bb7fe34caf09df5f22ffd9f661bc200bff93fcf56e15fb046a1125ec39bcd7e82a1fcec8026b2f0a204efc0ca16c863ae831ed26a5bd64e76bea6ee6c84aa00bdc77518a18b52f6f22b856487106b607823955b66715878fc788337edd40d69545db11270c3a6d202d018b6ea3a28d19d20628b7b82426cff4077e239f30c887363c5345cacfde6da0c0499b95c2bf1963a2e09681661a4ce2d173f1145cbadb42818af483967b4731e91a21d5d37423473e4d2919d591e0bffaba1562d7eff89b8cbf7bd25fa6f08c4663cdbe584a9e6f694e1088546dd2dc0677721048682d592287388c1f8e71c93cb10c7807fbcc6b7aaa11a1612b662921357e7aa0234e68194e3a5074bfd3134318ab86253dc7f43fe72fb48845da9d1f0d5bbec25e245c330cf072cd2cf70d4f6f7b7aa94698a13aeff015e8e8d9ed0e524d65dc6b29d1161ffe3513ba6247f69b7181086f61b2c64cfd5e23c8d95ff9652c2c873036056f22018b740bc8cd7dc680ad442d662fce7a90a6fcee047463012525622db50931c2c18a3fc369f90fdc59f84cb32dc07c10f00be4125c40be96e3c401fb8bacf28b37b1cc1424eddc3a773c885d76be2291153c408af2a8c94510bdba4b0b53941f59be812f6e9ccc655fc5efdcc6f71a9b148f96cea332902abbd12633cc03a97074e3d600ddfd684bdaabf6b909bc9d22c834215c21de89b9f5dbc84f6e64133009083842d029889efc9791b93715982c805c3a80dcfea831b56903392b99b06488bd0a465e0eac06946520f6d3917f2402ffe125b20e32f1d11b623192146157dfb55abf60acea565c387f53e776f75e8c54e11a2612e4a53f46c858ae6741f9dfa3"
    }
  ]
}`

//end of file