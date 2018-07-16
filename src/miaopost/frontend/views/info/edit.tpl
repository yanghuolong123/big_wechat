<div class="create">
        <form class="form-horizontal">
            <div class="form-group">            
              <div class="col-sm-offset-1 col-sm-4">
                <select id="cid" class="form-control">
                  <option value="">请选择分类</option>
                  {{range .cats}}
                    <option value="{{.Id}}" {{if  eq .Id  $.info.Cid}}selected{{end}}>{{.Name}}</option>
                  {{end}}
                </select>
              </div>
            </div>          
            <div class="form-group">             
              <div class="col-sm-offset-1 col-sm-6">
                <label>图片和文字可以只填一项</label>
                <div class="img-up">
                    <div class="img-up-list clearfix">                        
                      {{range .photos}}
                        <div class="img-li img-li-new" data-url="{{.Url}}"  data-big="' + upImg + '" style="background-image:url({{.Url}}!200!200)"><i></i></div>
                      {{end}}
                    </div>
                    <div id="thelist" class="uploader-list"></div>
                    <div id="picker"><label class="user-img" for="imgs"></label></div>
                </div>
                <textarea class="form-control"  id="info_content" rows="5" maxlength="2500" placeholder="添加描述">{{.info.Content}}</textarea>                
              </div>
            </div>
            <div class="form-group" style="{{if .user}}display: none;{{end}}">            
              <div class=" col-sm-offset-1 col-sm-2">
                <label>自动删除发布</label>
                <div class="row">
                  <div  class="col-md-8">
                  <input type="text" class="form-control"  id="valid_day" value="{{.info.Valid_day}}" maxlength="10" placeholder=""> 
                  </div>
                  <div class="col-md-4 valid_date">
                    <span class="">天后</span>
                  </div>
                </div>
              </div>
            </div>
             <div class="form-group" style="{{if .user}}display: none;{{end}}">            
              <div class=" col-sm-offset-1 col-sm-4">
                <label>我们会发送编辑链接至邮箱</label>
                <input type="text" class="form-control" id="email" value="{{.info.Email}}" maxlength="55" placeholder="邮箱地址">
              </div>
            </div>                  
            <div class="form-group">
                  <div class="col-sm-offset-1 col-sm-5  error_tips text-danger"></div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <input type="hidden" id="info_id"  value="{{.info.Id}}" name="id">
                <button type="button" id="edit_info_btn" class="btn btn-success btn-lg">保存</button>
                {{if not .user}}
                &nbsp;&nbsp;&nbsp;&nbsp;
                <button type="button" id="del_info_btn" class="btn btn-danger btn-lg">删除</button>
                {{end}}
              </div>
            </div>
          </form>
</div>
