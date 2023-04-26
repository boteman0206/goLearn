package git

/**
git的reset使用：
	是将master直接移动到当前commit上：有三种模式 soft，mixed， heard
	soft：回退后a分支修改的代码被保留并标记为add的状态（git status 是绿色的状态）
	mixed: 重置索引，但不重置工作树，更改后的文件标记为未提交（add）的状态。默认操作。
	heard： 重置索引和工作树，并且a分支修改的所有文件和中间的提交，没提交的代码都被丢弃了。

	执行完成之后需要强制推送到远程分支，好像master分支不能强制推送，需要设置，其他的分支可以强制推送
    执行完成之后原来的提交记录都不存在了
	git push -f 或者使用  git push -f origin v1.0 (git push -f 远程仓库名 当前分支名:远程其他分支名)

git revert使用一个新的提交来更新代码，原来的提交记录都保留


区别：
git revert是用一次新的commit来回滚之前的commit，git reset是直接删除指定的commit。



one


two

three
*/
