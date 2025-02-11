<template>
  <div :style="{ padding: '20px 20px 12px 20px' }">
    <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
      <a-space>
        <a-input v-model:value="query.id" placeholder="合同编号" style="width: 250px; margin-right: 10px;">
          <template #suffix>
            <search-outlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch"/>
          </template>
        </a-input>
        <a-button :type="buttonType.bt1" @click="onContractList">全部合同</a-button>
        <a-button :type="buttonType.bt2" @click="onContractStatus(1)">已签约合同</a-button>
        <a-button :type="buttonType.bt3" @click="onContractStatus(2)">未签约合同</a-button>
        <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
      </a-space>
      <div>
        <a-space size="middle">
          <a-button type="primary" @click="onCreate">新建</a-button>
          <a-button type="primary" @click="onExport" ghost>
            <template #icon>
              <ExportOutlined/>
            </template>
            导出
          </a-button>
        </a-space>
      </div>
    </div>
    <a-table rowKey="id" :rowSelection="{ selectedRowKeys: tableData.selectedIds, onChange: onSelectChange }"
             :columns="columns" :dataSource="data.contractList"
             :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPagination }"
             :scroll="{ y: '59vh' }" class="ant-table-striped"
             :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered
             @resizeColumn="handleResizeColumn"
    >
      <template #bodyCell="{ column, text, record }">
        <template v-if="column.dataIndex === 'id' || column.dataIndex === 'name'">
          <a @click="onEdit(record)">{{ text }}</a>
        </template>
        <template v-if="column.dataIndex === 'amount'">
          <span style="color: #ff991f">{{ text }}</span>
        </template>
        <template v-if="column.dataIndex === 'status'">
          <Spot v-if="text === 1" type="success" title="已签约"/>
          <Spot v-if="text === 2" type="danger" title="未签约"/>
        </template>
      </template>
    </a-table>
    <!-- 新建、编辑合同 -->
    <a-modal v-model:open="open" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
             width="800px" :centered="true">
      <div style="height: 55vh;overflow-y: scroll;padding: 0 15px;">
        <a-form ref="contractFormRef" :model="contract" layout="vertical" name="contract" :rules="rules">
          <a-row :gutter="16">
            <a-col :span="12">
              <a-form-item label="合同编号" name="id">
                <a-input v-model:value="contract.id" :disabled="true" placeholder="根据编号规则自动生成"/>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="合同名称" name="name">
                <a-input v-model:value="contract.name"/>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="16">
            <a-col :span="12">
              <a-form-item label="客户名称" name="cid">
                <a-select v-model:value="contract.cid" style="width: 100%;" placeholder="请选择"
                          :fieldNames="{ label: 'name', value: 'id' }" :options="data.customerOption"
                          @focus="getCustomerOption"
                          @change="changeCustomerOption"></a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="合同金额" name="amount">
                <a-input v-model:value="contract.amount" style="width: 100%" :disabled="true"/>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="16">
            <a-col :span="12">
              <a-form-item label="合同开始时间" name="beginTime">
                <a-date-picker v-model:value="contract.beginTime" placeholder="选择日期" style="width: 100%"
                               format="YYYY-MM-DD" valueFormat="YYYY-MM-DD"/>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="合同结束时间" name="overTime">
                <a-date-picker v-model:value="contract.overTime" placeholder="选择日期" style="width: 100%"
                               format="YYYY-MM-DD" valueFormat="YYYY-MM-DD"/>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="16">
            <a-col :span="12">
              <a-form-item label="合同状态" name="status">
                <a-select v-model:value="contract.status" placeholder="请选择">
                  <a-select-option :value="1">已生效</a-select-option>
                  <a-select-option :value="2">未生效</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="备注" name="remarks">
                <a-input v-model:value="contract.remarks"/>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row :gutter="16">
            <a-col :span="24">
              <a-form-item label="产品" name="product">
                <a-button type="primary" @click="onAddProduct" style="float: right;margin-bottom: 10px;">
                  添加产品
                </a-button>
                <a-table rowKey="id" :columns="productColumns" :data-source="data.addedProductList"
                         :scroll="{ y: '59vh' }" class="ant-table-striped"
                         :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)"
                         :pagination="false"
                         bordered>
                  <template #bodyCell="{ column, text, record }">
                    <template v-if="column.dataIndex === 'type'">
                      <span v-if="text === 1">默认</span>
                    </template>
                    <template v-if="column.dataIndex === 'price'">
                      <span style="color: #ff991f">{{ text }}</span>
                    </template>
                    <template v-if="column.dataIndex === 'count'">
                      <!--使用a-form-item-rest解决3.x问题-->
                      <a-form-item-rest>
                        <a-input-number v-model:value="record.count" @change="calculatedAmount"/>
                      </a-form-item-rest>
                    </template>
                    <template v-if="column.dataIndex === 'total'">
                      <span>{{ record.total = record.price * record.count }}</span>
                    </template>
                    <template v-if="column.dataIndex === 'operation'">
                      <a @click="delProduct(record)">删除</a>
                    </template>
                  </template>
                </a-table>
              </a-form-item>
            </a-col>
          </a-row>
          <a-row>
            <a-col :span="24">
              <div style="float: right;margin: 0 20px;">
                <span>已选择产品：{{ data.addedProductList.length }}种&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;总金额：</span>
                <a-input-number v-model:value="contract.amount" style="width: 200px;"/>
              </div>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </a-modal>
    <!-- 添加产品 -->
    <a-modal v-model:open="productListVisible" title="添加产品" @cancel="onCancelProductList" @ok="onConfirm"
             cancelText="取消" okText="确定" width="800px" style="top: 80px">
      <div style="height: 55vh;padding: 0 15px;">
        <a-table rowKey="id"
                 :row-selection="{ selectedRowKeys: data.defaultSelectedIds, onChange: onSelectedProductIds }"
                 :columns="productListColumns" :data-source="data.productList"
                 :pagination="{ current: pagination.current, pageSize: pagination.pageSize, total: pagination.total, onChange: onPaginationProduct }"
                 :scroll="{ y: '40vh' }" class="ant-table-striped"
                 :row-class-name="(_record, index) => (index % 2 === 1 ? 'table-striped' : null)" bordered>
          <template #bodyCell="{ column, text, record }">
            <template v-if="column.dataIndex === 'name'">
              <a @click="onEdit(record)">{{ text }}</a>
            </template>
            <template v-if="column.dataIndex === 'status'">
              <span v-if="text === 1">上架</span>
              <span v-if="text === 2">下架</span>
            </template>
            <template v-if="column.dataIndex === 'type'">
              <span v-if="text === 1">默认</span>
            </template>
            <template v-if="column.dataIndex === 'price'">
              <span style="color: #ff991f">{{ text }}</span>
            </template>
          </template>
        </a-table>
      </div>
    </a-modal>
  </div>
</template>

<script setup>
import {createVNode, defineAsyncComponent, onMounted, reactive, ref} from 'vue';
import {ExclamationCircleOutlined, ExportOutlined, SearchOutlined} from '@ant-design/icons-vue';
import moment from 'moment'
import {
  contractExport,
  createContract,
  deleteContract,
  queryContractInfo,
  queryContractList,
  queryContractPlist,
  updateContract
} from '../api/contract';
import {queryProductList} from "../api/product";
import {queryCustomerOption} from "../api/customer";
import {message, Modal} from 'ant-design-vue';
import Spot from '../components/Spot.vue';

// 条件筛选
const query = reactive({
  id: undefined,
  status: undefined
})


const tableData = reactive({
  selectedIds: []
})

// 点击搜索合同
const onSearch = () => {
  contractList()
}

const defaultSelected = (record) => {
  return {
    disabled: true,
    name: record
  }
}

// 点击全部合同
const onContractList = () => {
  query.id = undefined
  query.status = undefined
  pagination.current = 1
  setButtonType()
  contractList()
}

// 点击已签约或未签约合同
const onContractStatus = (status) => {
  query.status = status
  pagination.current = 1
  setButtonType(status)
  contractList()
}

// 按钮默认类型
const buttonType = reactive({
  bt1: 'primary',
  bt2: 'default',
  bt3: 'default',
})

// 设置按钮类型
const setButtonType = (status) => {
  switch (status) {
    case 1:
      buttonType.bt1 = 'default'
      buttonType.bt2 = 'primary'
      buttonType.bt3 = 'default'
      break;
    case 2:
      buttonType.bt1 = 'default'
      buttonType.bt2 = 'default'
      buttonType.bt3 = 'primary'
      break;
    default:
      buttonType.bt1 = 'primary'
      buttonType.bt2 = 'default'
      buttonType.bt3 = 'default'
      break;
  }
}

// 合同表格字段
const columns = ref([
  {
    title: '合同编号',
    dataIndex: 'id',
    width: 50,
    fixed: 'left',
    ellipsis: true,
    resizable: true,
    minWidth: 50,
    maxWidth: 100,
  }, {
    title: '合同名称',
    dataIndex: 'name',
    width: 200,
    fixed: 'left',
    ellipsis: true,
    resizable: true,
    minWidth: 200,
    maxWidth: 300,
  }, {
    title: '客户名称',
    dataIndex: 'cname',
    width: 100,
    resizable: true,
    minWidth: 100,
    maxWidth: 200,
  }, {
    title: '合同金额',
    dataIndex: 'amount',
    width: 100,
  }, {
    title: '合同开始时间',
    dataIndex: 'beginTime',
    width: 150,
  }, {
    title: '合同结束时间',
    dataIndex: 'overTime',
    width: 150,
  }, {
    title: '备注',
    dataIndex: 'remarks',
    width: 240,
    ellipsis: true,
  }, {
    title: '签约状态',
    dataIndex: 'status',
    width: 100,
    ellipsis: true,
  }, {
    title: '创建时间',
    dataIndex: 'created',
    width: 185,
    customRender: text => {
      return text.value === 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
  }, {
    title: '更新时间',
    dataIndex: 'updated',
    width: 185,
    customRender: text => {
      return text.value === 0 ? '' : moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
    }
  }]);

// 新建或编辑合同，已添加产品表格字段
const productColumns = [
  {
    title: '产品名称',
    dataIndex: 'name',
    width: 100,
  }, {
    title: '产品类别',
    dataIndex: 'type',
    width: 90,
  }, {
    title: '单位',
    dataIndex: 'unit',
    width: 80,
  }, {
    title: '价格',
    dataIndex: 'price',
    width: 100,
  }, {
    title: '数量',
    dataIndex: 'count',
    width: 120,
  }, {
    title: '合计',
    dataIndex: 'total',
    width: 100,
  }, {
    title: '操作',
    dataIndex: 'operation',
    width: 100,
  }]

// 产品表格字段
const productListColumns = [
  {
    title: '产品名称',
    dataIndex: 'name',
    width: 100,
    fixed: 'left',
    ellipsis: true,
  }, {
    title: '是否上下架',
    dataIndex: 'status',
    width: 120,
  }, {
    title: '产品类型',
    dataIndex: 'type',
    width: 100,
  }, {
    title: '产品单位',
    dataIndex: 'unit',
    width: 100,
  }, {
    title: '产品编码',
    dataIndex: 'code',
    width: 150,
  }, {
    title: '价格',
    dataIndex: 'price',
    width: 150,
  }, {
    title: '产品描述',
    dataIndex: 'description',
    width: 240,
    ellipsis: true,
  }, {
    title: '创建时间',
    dataIndex: 'created',
    width: 185,
    customRender: text => {
      let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
      return m === 'Invalid date' ? '' : m
    }
  }, {
    title: '更新时间',
    dataIndex: 'updated',
    width: 185,
    customRender: text => {
      let m = moment(text.value * 1000).format('YYYY-MM-DD HH:mm:ss')
      return m === 'Invalid date' ? '' : m
    }
  }, {
    title: '创建人',
    dataIndex: 'creator',
    width: 150,
  }];

// 表单校验
const rules = {
  name: [{required: true, message: '请输入合同名称', trigger: 'blur'}],
  cid: [{required: true, message: '请选择客户', trigger: 'blur'}],
  status: [{required: true, message: '请选择合同状态'}]
};

// 合同属性
let contract = reactive({
  id: undefined,
  name: undefined,
  amount: undefined,
  beginTime: '',
  overTime: '',
  cid: undefined,
  remarks: undefined,
  status: undefined,
  productList: [],
});

const data = reactive({
  contractId: 0,
  contractList: [],
  contractIds: [],
  productList: [],
  productIds: [],
  addedProductList: [],
  customerOption: [],
  defaultSelectedIds: []
})

// 表格分页
let pagination = reactive({
  current: 1,
  pageSize: 10,
  total: undefined,
})

const title = ref('');
const open = ref(false);
const disabled = ref(true)
const operation = ref(0);
const contractFormRef = ref();
const productListVisible = ref(false);

// 点击新建合同
const onCreate = () => {
  title.value = '新建合同'
  operation.value = 1
  open.value = true
}

// 点击编辑合同
const onEdit = (row) => {
  title.value = '编辑合同'
  operation.value = 2
  getCustomerOption()
  let param = {id: row.id}
  queryContractInfo(param).then((res) => {
    if (res.data.code === 0) {
      let p = res.data.data
      contract.id = p.id
      contract.name = p.name
      contract.cid = p.cid
      contract.amount = p.amount
      contract.beginTime = p.beginTime
      contract.overTime = p.overTime
      contract.remarks = p.remarks
      contract.status = p.status
      contract.productList = p.product_list || []
      data.addedProductList = p.product_list || []
    }
  })
  data.contractId = row.id
  open.value = true
}

// 点击保存合同
const onSave = () => {
  if(data.addedProductList.length === 0){
    message.info('最少选择应该产品')
    return
  }

  contractFormRef.value.validateFields().then(() => {
    if (operation.value === 1) {
      let param = {
        name: contract.name,
        cid: contract.cid,
        amount: contract.amount,
        beginTime: contract.beginTime,
        overTime: contract.overTime,
        remarks: contract.remarks,
        status: contract.status,
        product_list: data.addedProductList,
      }
      createContract(param).then((res) => {
        if (res.data.code === 0) {
          message.success('保存成功')
          data.defaultSelectedIds = []
          contractList()
        }
      })
    }
    if (operation.value === 2) {
      let param = {
        id: contract.id,
        name: contract.name,
        cid: contract.cid,
        amount: contract.amount,
        beginTime: contract.beginTime,
        overTime: contract.overTime,
        remarks: contract.remarks,
        status: contract.status,
        product_list: data.addedProductList,
      }
      updateContract(param).then((res) => {
        if (res.data.code === 0) {
          message.success('保存成功')
          data.defaultSelectedIds = []
          contractList()
        }
      })
    }
    contractFormRef.value.resetFields()
    open.value = false;
  });
};

// 点击删除合同
const onDelete = () => {
  let param = {
    ids: tableData.selectedIds
  }
  Modal.confirm({
    title: '确定删除选中的' + tableData.selectedIds.length + '项吗?',
    icon: createVNode(ExclamationCircleOutlined),
    centered: true,
    cancelText: '取消',
    okText: '确定',
    onOk() {
      deleteContract(param).then((res) => {
        if (res.data.code === 0) {
          contractList()
          disabled.value = true
          message.success('删除成功')
        }
      })
    },
    onCancel() {
    },
  });
}

// 处理表头自适应宽度
const handleResizeColumn = (w, col) =>{
  col.width = w;
}

// 初始化数据
onMounted(() => {
  contractList()
})

// 分页查询合同列表
const onPagination = (page) => {
  pagination.current = page
  contractList(query.status)
}
const contractList = () => {
  let param = {
    id: parseInt(query.id === undefined || query.id === '' ? '0' : query.id),
    status: query.status,
    pageNum: pagination.current,
    pageSize: pagination.pageSize
  }
  queryContractList(param).then((res) => {
    if (res.data.code === 0) {
      pagination.total = res.data.data.total
      data.contractList = res.data.data.list
    }
  })
}

// 点击添加产品
const onAddProduct = () => {
  let param = {
    pageNum: pagination.current,
    pageSize: pagination.pageSize
  }
  queryProductList(param).then((res) => {
    if (res.data.code === 0) {
      pagination.total = res.data.data.total
      data.productList = res.data.data.list
    }
  })
  data.defaultSelectedIds = []
  if (data.addedProductList.length > 0) {
    for (let i = 0; i < data.addedProductList.length; i++) {
      data.defaultSelectedIds[i] = data.addedProductList[i].id
    }
  }
  productListVisible.value = true
}

// 分页查询产品列表
const onPaginationProduct = (page) => {
  pagination.current = page
  let param = {
    pageNum: pagination.current,
    pageSize: pagination.pageSize
  }
  queryProductList(param).then((res) => {
    if (res.data.code === 0) {
      pagination.total = res.data.data.total
      data.productList = res.data.data.list
    }
  })
}

// 已选中的合同ID
const onSelectChange = (selectedRowKeys) => {
  tableData.selectedIds = selectedRowKeys
  disabled.value = tableData.selectedIds.length === 0;
};

// 已选中的产品ID
const onSelectedProductIds = selectedRowKeys => {
  data.productIds = selectedRowKeys
  data.defaultSelectedIds = selectedRowKeys
};

// 删除选中的产品
const delProduct = (row) => {
  for (let i = 0; i < data.addedProductList.length; i++) {
    if (data.addedProductList[i].id === row.id) {
      data.addedProductList.splice(i, 1);
    }
  }
  calculatedAmount()
}

// 点击确定选中的产品ID
const onConfirm = () => {
  let param = {
    id: data.contractId,
    pids: data.productIds
  }
  queryContractPlist(param).then((res) => {
    if (res.data.code === 0) {
      data.addedProductList = res.data.data
    }
  })
  productListVisible.value = false
}

// 查询客户选项
const getCustomerOption = () => {
  queryCustomerOption().then((res) => {
    if (res.data.code === 0) {
      data.customerOption = res.data.data
    }
  })
}

const changeCustomerOption = (value) => {
  contract.cid = value
}

// 计算金额
const calculatedAmount = () => {
  contract.amount = 0
  let totalAmount = 0
  for (let i = 0; i < data.addedProductList.length; i++) {
    totalAmount = totalAmount + (data.addedProductList[i].price * data.addedProductList[i].count)
  }
  contract.amount = totalAmount
}

// 点击合同取消按钮
const onCancel = () => {
  contractFormRef.value.resetFields()
  data.addedProductList = []
  data.contractId = undefined
  open.value = false
};

// 导出表格
const onExport = () => {
  contractExport().then((res) => {
    if (res.data.type === 'application/json') {
      message.error('导出错误！')
    } else {
      let blob = new Blob([res.data], {
        type: "application/vnd.ms-excel"
      })
      let a = document.createElement('a')
      a.setAttribute("download", "合同信息.xlsx");
      a.href = window.URL.createObjectURL(blob)
      a.click()
      window.URL.revokeObjectURL(a.href)
    }
  })
}

// 点击取消产品列表
const onCancelProductList = () => {
  productListVisible.value = false
  data.contractId = undefined
  pagination.current = 1
  pagination.total = undefined
}
</script>

<style scoped>
.ant-table-striped :deep(.table-striped) td {
  background-color: #fafafa;
}
</style>