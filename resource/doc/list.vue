<h3>文章列表</h3>
<hr/>
<div id="app"></div>
<template id="template">
    <table class="table table-striped  table-bordered">
        <thead>
        <tr>
            <th>N.O</th>
            <th>城市</th>
            <th>标签</th>
            <th>创建时间</th>
            <th>更新时间</th>
        </tr>
        <tbody>
        <tr v-for="(city, i) in list">
            <td>{{city.id}}</td>
            <td><a href="javascript:;" @click="greet(city.title, city.content)">{{city.title}}</a></td>
            <td>{{city.tag}}</td>
            <td>{{city.create_time}}</td>
            <td>{{city.update_time}}</td>
        </tr>
    </table>

    <div class="modal fade" id="share" tabindex="-1" role="dialog">
        <div class="modal-dialog" role="document" style="width: 80%;">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                            aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title">{{title}}</h4>
                </div>
                <div class=" modal-body">
                    <div id="markdown-view">
                    </div>
                </div>

                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">确认</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script type="text/javascript">
$(document).ready(function () {
    var vm = Vue.createApp({
        data() {
            let list = JSON.parse(__JS_DATA__.api_data)
            return {
                list: list,
                title : '',
                content: '',
            }
        },
        template: '#template',
        methods: {
            async greet(title, content) {
                this.title = title                
                $('#share').modal({
                    keyboard:false,
                    backdrop : 'static',
                })
                $('#markdown-view').html('')
                editormd.markdownToHTML("markdown-view", {
                    "markdown" : content,
                });
                
            }
        }
    })
    vm.mount('#app')
})
</script>