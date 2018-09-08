<div class="user_menu">
        <a {{if  eq .uri "/user/my"}}class="label label-primary"{{end}} href="/user/my">我的发布</a>
        <a {{if  eq .uri "/user/edit"}}class="label label-primary"{{end}} href="/user/edit">修改昵称</a>
        <a {{if  eq .uri "/user/account"}}class="label label-primary"{{end}} href="/user/account">我的钱包</a>
        <a {{if  eq .uri "/user/adv"}}class="label label-primary"{{end}} href="/user/adv">我的广告</a>
        {{if not .isWeixin}}            
        <a href="/logout" class=""><span class="glyphicon glyphicon-log-out"></span> 退出</a>
        {{end}}
</div>
