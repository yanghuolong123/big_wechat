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
                <label>图片和文字 <span class="text-danger label-tips">(至少只填一项)</span></label>
                <textarea class="form-control"  id="info_content" rows="3" maxlength="2500" placeholder="添加描述"></textarea>
                <div class="img-up">
                    <div class="img-up-list clearfix">                        
                      
                    </div>
                    <label class="user-img" for="imgs"><input type="file" name="imgs" id="imgs"></label>
                </div>
              </div>
            </div>
            <div class="form-group">            
              <div class=" col-sm-offset-1 col-sm-2">
                <label>自动删除发布 <span class="text-danger label-tips">(建议填写)</span></label>
                <div class="row">
                  <div  class="col-md-8">
                  <input type="text" class="form-control"  id="valid_day" maxlength="10" placeholder=""> 
                  </div>
                  <div class="col-md-4 valid_date">
                    <span class="">天后</span>
                  </div>
                </div>
              </div>
            </div>
             <div class="form-group">            
              <div class=" col-sm-offset-1 col-sm-4">
                <label>我们会发送编辑链接至邮箱 <span class="text-danger label-tips">(建议填写)</span></label>
                <input type="text" class="form-control" id="email" maxlength="55" placeholder="邮箱地址">
              </div>
            </div>                  
            <div class="form-group">
                  <div class="col-sm-offset-1 col-sm-5  error_tips text-danger"></div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <button type="button" id="create_info_btn" class="btn btn-primary btn-lg">发布</button>
              </div>
            </div>
          </form>
</div>
