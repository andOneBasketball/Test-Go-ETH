## 编译Solidity文件
solc编译bin和abi 文件
solc --abi --bin test.sol 
若是想要输出文件则是：
solc --abi --bin -o ./  test.sol
abigen 结合bin和abi编译Golang文件
abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go