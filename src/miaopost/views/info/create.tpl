<div class="create">
        <form class="form-horizontal">
            <div class="form-group">            
              <div class="col-sm-offset-1 col-sm-4">
                <select id="cid" class="form-control">
                  <option value="">请选择分类</option>
                  {{range .cats}}
                    <option value="{{.Id}}">{{.Name}}</option>
                  {{end}}
                </select>
              </div>
            </div>          
            <div class="form-group">             
              <div class="col-sm-offset-1 col-sm-6">
                <label>图片和文字可以只填一项</label>
                <textarea class="form-control"  id="info_content" rows="3" maxlength="250" placeholder="添加描述"></textarea>
              </div>
            </div>
            <div class="form-group">            
              <div class=" col-sm-offset-1 col-sm-2">
                <label>自动删除发布</label>
                <input type="text" class="form-control"  id="valid_day" maxlength="10" placeholder=""> 
              </div>
              <div class="col-sm-2 valid_date"><span class="">天后</span></div>
            </div>
             <div class="form-group">            
              <div class=" col-sm-offset-1 col-sm-4">
                <label>我们会发送编辑链接至邮箱</label>
                <input type="text" class="form-control" id="email" maxlength="55" placeholder="邮箱地址">
              </div>
            </div>                  
            <div class="form-group">
                  <div class="col-sm-offset-1 col-sm-5  error_tips text-danger"></div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <button type="button" id="create_info_btn" class="btn btn-success btn-lg">发布</button>
              </div>
            </div>
          </form>
</div>
