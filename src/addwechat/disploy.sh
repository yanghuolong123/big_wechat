#!/bin/bash

HOST="104.192.6.54"
USERNAME="root"
PORT="22"
DIR="/var/work/web/goweb/addwechat"
#HOST="219.234.0.171"
#USERNAME="root"
#PORT="22000"
#DIR="/home/work/www/big_wechat"

echo $USERNAME@$HOST:$DIR

ssh -p $PORT $USERNAME@$HOST > /dev/null 2>&1 << eeooff
killall addwechat
rm -rf $DIR/addwechat
rm -rf $DIR/conf
rm -rf $DIR/static
rm -rf $DIR/views
exit
eeooff
echo "停止旧有项目!"
echo "删除旧有目录完成!"

scp -P$PORT addwechat $USERNAME@$HOST:$DIR
scp -P$PORT -r conf $USERNAME@$HOST:$DIR
scp -P$PORT -r static $USERNAME@$HOST:$DIR
scp -P$PORT -r views $USERNAME@$HOST:$DIR

echo "部署完成!"

ssh -p $PORT $USERNAME@$HOST > /dev/null 2>&1 << eeooff
cd $DIR
#sed -i 's/mysql.user = root/mysql.user = feichangjuzu/g' ./conf/app.conf
sed -i 's/mysql.pass = 123456/mysql.pass = Dn1h345RIC/g' ./conf/app.conf
sed -i 's/wechat.token = feichangjuzu123456/wechat.token = big_wechat123456/g' ./conf/app.conf
sed -i 's/wechat.appid = wx2705fb0b58b923b6/wechat.appid = wx3e0b8bca5b6d6606/g' ./conf/app.conf
sed -i 's/wechat.secret = 63b572bc483358797be65ea66b156290/wechat.secret = 3eaf2d4ce9c9c36303c087020e3e04cc/g' ./conf/app.conf
./addwechat > ./log/web.log &
exit    
eeooff

echo "项目启动成功!"

