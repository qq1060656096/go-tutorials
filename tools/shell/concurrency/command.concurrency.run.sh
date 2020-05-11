#!/usr/bin/env bash
############################
# github: https://github.com/qq1060656096
# author: zwei
# mail: 1060656096@qq.com
# name: 并发测试脚本
# version: 1.0
# description: 并发测试命令,支持多个命令
#
############################
appName="并发测试脚本"
appVersion="1.0"
appDir=`pwd`

: << !
并发运行命令
concurrentRunCommand 运行次数 运行命令 命令日志文件
concurrentRunCommand 10 "echo 1" "echolog"
!
function concurrentRunCommand() {
    # 运行次数
    runNumber=$1
    # 命令
    runCommand=$2
    # 命令日志
    commandLog=$3
    for (( i = 1; i <= ${runNumber}; i++ ));
    do
        `${runCommand} >> ${commandLog}.${i}.log`&
    done
   
}

if [ ${#} -lt 3 ]
then
    echo "############################"
    echo "# github: https://github.com/qq1060656096"
    echo "# author: zwei"
    echo "# mail: 1060656096@qq.com"
    echo "# name: 并发测试脚本"
    echo "# version: 1.0"
    echo "# description: 并发测试命令,支持多个命令"
    echo "# useage: command.concurrency.run.sh "
    echo "#     command.concurrency.run.sh runNumber commandFile commandLog"
    echo "#     command.concurrency.run.sh 运行次数 命令文件 日志文件"
    echo "############################"
    echo ""
    echo -e "\033[38;5;1;4m参数出错误\033[0m"
    exit   
fi
echo ${shellFile} ${runNumber} ${commandLog}


runNumber=${1}
shellFile=${2}
commandLog=${3}
concurrentRunCommand ${runNumber} ${shellFile} ${commandLog}
