作业说明

> Q：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

> A：dao层应该将sql.ErrNoRows往上抛，1、如果不抛给上层，一般做法是通过判定后返回nil给上层，每一层都要进行一次判定处理，2、另外如果上层需要处理sql.ErrNoRows，来返回不同错误码给调用方时是无法判定的，3、如果不通过wrap方式抛error，无法获取错误堆栈信息，无法清晰定位问题。

> 代码地址：https://github.com/InjuryOf/Go-000/blob/main/Week02/work.go

