<script src="/static/plugin/jquery/jquery-2.2.4.js"></script>
<script src="/static/plugin/jquery.mobile/jquery.mobile-1.4.5.min.js"></script>
<script src="/static/plugin/bootstrap/js/bootstrap.min.js"></script>
<script>
var sessionId={{.user.Id}};
var gid={{.user.Gid}};
var follow={{.follow}};
var nickname={{.user.Nickname}}
var groupname={{.group.Name}}

$(function(){
	$('input[type="text"],textarea').on('click', function () {
  		var target = this;
  		setTimeout(function(){
        		target.scrollIntoViewIfNeeded();
      		},400);
	});
});
</script>
<script src="/static/js/main.js"></script>
