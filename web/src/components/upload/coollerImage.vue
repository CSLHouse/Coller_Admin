
<template>
    <div>
      <el-upload
        :list-type="listType"
        :action="`${path}/fileUploadAndDownload/upload`"
        :headers="{ 'x-token': userStore.token }"
        :show-file-list="isShow"
        :on-success="handleImageSuccess"
        :before-upload="beforeImageUpload"
        :on-remove="handleRemove"
        :on-preview="handlePreview"
        :limit="maxCount"
        multiple
        ref="uploadRef"
        :file-list="fileDetailList"
      >
      <!-- :auto-upload="false" -->
        <el-icon><Plus /></el-icon>
        <template #tip>
            <div class="el-upload__tip">
                只能上传jpg/png文件，且不超过10MB
            </div>
        </template>
      </el-upload>
    </div>
  </template>
  
  <script setup>
  import ImageCompress from '@/utils/image'
  import { ref, toRefs, defineExpose, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useUserStore } from '@/pinia/modules/user'
  import { deleteFile } from '@/api/fileUploadAndDownload'

  const emit = defineEmits(['change'])
  const props = defineProps({
    imageUrl: {
      type: String,
      default: ''
    },
    fileSize: {
      type: Number,
      default: 2048 // 2M 超出后执行压缩
    },
    maxWH: {
      type: Number,
      default: 1920 // 图片长宽上限
    },
    maxCount: {
      type: Number,
      default: 5 // 最多5张
    },

    isShow: {
      type: Boolean,
      default: true
    },
    listType: {
      type: String,
      default: "picture-card"
    },
    fileList: Array
  })

  const path = ref(import.meta.env.VITE_BASE_API)
  const userStore = useUserStore()
  const fileDetailList = toRefs(props.fileList)

  computed: {
      for(let i=0; i< fileDetailList.length; i++){
        fileDetailList[i] = {url: props.fileList[i]}
      }
  }

  const beforeImageUpload = (file) => {
    const isJPG = file.type === 'image/jpeg'
    const isPng = file.type === 'image/png'
    if (!isJPG && !isPng) {
      ElMessage.error('上传头像图片只能是 jpg或png 格式!')
      return false
    }
  
    const isRightSize = file.size / 1024 < props.fileSize
    if (!isRightSize) {
      // 压缩
      const compress = new ImageCompress(file, props.fileSize, props.maxWH)
      return compress.compress()
    }
    return isRightSize
  }
  
  const handleImageSuccess = (res) => {
    const { data } = res
    if (data.file) {
      fileDetailList.push({url: data.file.url})
      // console.log("==[handleImageSuccess]===fileDetailList:", fileDetailList)
      emit('change', fileDetailList)
    }
  }
  
  const handleRemove = async(res) => {
    console.log("-----res-----", res)
    const { response } = res
    if (response.data) {
      const res = await deleteFile(response.data.file)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!',
        })
        emit('change', fileDetailList)
        console.log("----[handleRemove]------fileList:", fileDetailList)
      }
    }
  }
  const handleError = () => {
    console.log("----[handleError]------上传失败:")
    ElMessage.error("上传失败!");
  }
  const handlePreview = (url) => {
    ElMessage.warn('已上传过同名文件!')
    return
  }

  const uploadRef = ref()
  const submitUpload = () => {
    uploadRef.value.submit()
  }

  defineExpose({
    submitUpload,
  })
  </script>
  
  <script>
  
  export default {
    name: 'UploadImage',
    methods: {
  
    }
  }
  </script>
  
  <style lang="scss" scoped>
  .image-uploader {
    border: 1px dashed #d9d9d9;
    width: 180px;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .image-uploader {
    border-color: #409eff;
  }
  .image-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .image {
    width: 178px;
    height: 178px;
    display: block;
  }

  .upload-queue{
  .el-upload-list__item:hover{
    .el-upload-list__item-actions{
      display: block;
      opacity: 1;
    }
  }

  .el-upload--picture-card .el-upload--text{
    width: 100%;
  }
}
  </style>
  