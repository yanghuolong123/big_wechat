<div class="entry">
	<div class="jumbotron">
	  <p class="text-success">请选择学校：(下次会自动进入您选定学区)</p>
	  <p>
	  {{range .regions}}
	  	<a class="btn btn-success" href="/setRegion?rid={{.Id}}">{{.Fullname}}</a>
	  {{end}}
	  </p>
	</div>
</div>