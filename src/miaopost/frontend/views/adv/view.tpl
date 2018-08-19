<div class="article-view row">

{{str2html .adv.A.Content}}

</div>

<div class="row">
	<div class="col-md-8">
		{{str2html .adv.A.Content}}
	</div>
	<div class="col-md-8 photos">
		{{range  .adv.Photos}}
		<img src="{{.}}">
		{{end}}
	</div>
</div>