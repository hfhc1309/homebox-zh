<template>
  <div>
    <AppImportDialog v-model="modals.import" />
    <BaseContainer class="flex flex-col gap-4 mb-6">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <MdiFileChart class="mr-2" />
            <span> 文件生成 </span>
            <template #description> 为您的库存生成不同的文件。 </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="navigateTo('/reports/label-generator')">
            <template #title>资产 ID 标签</template>
            为一系列资产 ID 生成可打印的 PDF 标签。这些标签并非特定于您的库存，因此您可以提前打印标签并在收到标签时将其贴到您的库存上。
            <template #button>
              标签生成
              <MdiArrowRight class="ml-2" />
            </template>
          </DetailAction>
          <DetailAction @action="getBillOfMaterials()">
            <template #title>物料清单</template>
            生成可导入电子表格程序的 TSV（制表符分隔值）文件。这是您的库存摘要，包含基本商品和定价信息。
            <template #button> 生成BOM </template>
          </DetailAction>
        </div>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <MdiDatabase class="mr-2" />
            <span> 导入 / 导出 </span>
            <template #description>
              将您的库存导入和导出到 CSV 文件。这对于将您的库存迁移到 Homebox 的新实例非常有用。
            </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="modals.import = true">
            <template #title>导入库存</template>
            导入 Homebox 的标准 CSV 格式。这<b>不会</b>覆盖库存中的任何现有物品。它只会添加新物品。
          </DetailAction>
          <DetailAction @action="getExportTSV()">
            <template #title>导出库存</template>
            导出 Homebox 的标准 CSV 格式。这将导出库存中的所有物品。
          </DetailAction>
        </div>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <MdiAlert class="mr-2" />
            <span> 库存操作 </span>
            <template #description>
              批量对库存应用操作。这些操作不可逆转。<b>请小心。</b>
            </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="ensureAssetIDs">
            <template #title>确保资产 ID</template>
            确保库存中的所有物品都具有有效的 asset_id 字段。这是通过查找数据库中当前最高的 asset_id 字段并将下一个值应用于每个具有未设置 asset_id 字段的物品来实现的。这是按照 created_at 字段的顺序完成的。
          </DetailAction>
          <DetailAction @action="ensureImportRefs">
            <template #title>确保导入引用</template>
            确保库存中的所有物品都具有有效的 import_ref 字段。这是通过为每个未设置 import_ref 字段的物品随机生成一个 8 个字符的字符串来实现的。
          </DetailAction>
          <DetailAction @action="resetItemDateTimes">
            <template #title> 零项目日期时间</template>
            将清单中所有日期时间字段的时间值重置为日期的开始时间。这是为了修复网站开发初期引入的一个错误，该错误导致时间值与时间一起存储，从而导致日期字段无法显示准确值。
            <a class="link" href="https://github.com/hay-kot/homebox/issues/236" target="_blank">
              有关更多详细信息，请参阅 Github 问题 #236。
            </a>
          </DetailAction>
          <DetailAction @action="setPrimaryPhotos">
            <template #title> 设置主要照片 </template>
            在 Homebox v0.10.0 版本中，照片类型的附件中添加了主图像字段。如果尚未设置，此操作会将主图像字段设置为数据库中附件数组中的第一个图像。
            <a class="link" href="https://github.com/hay-kot/homebox/pull/576">请参阅 GitHub PR #576</a>
          </DetailAction>
        </div>
      </BaseCard>
    </BaseContainer>
  </div>
</template>

<script setup lang="ts">
  import MdiFileChart from "~icons/mdi/file-chart";
  import MdiArrowRight from "~icons/mdi/arrow-right";
  import MdiDatabase from "~icons/mdi/database";
  import MdiAlert from "~icons/mdi/alert";

  definePageMeta({
    middleware: ["auth"],
  });
  useHead({
    title: "Homebox | Profile",
  });

  const modals = ref({
    item: false,
    location: false,
    label: false,
    import: false,
  });

  const api = useUserApi();
  const confirm = useConfirm();
  const notify = useNotifier();

  function getBillOfMaterials() {
    const url = api.reports.billOfMaterialsURL();
    window.open(url, "_blank");
  }

  function getExportTSV() {
    const url = api.items.exportURL();
    window.open(url, "_blank");
  }

  async function ensureAssetIDs() {
    const { isCanceled } = await confirm.open(
      "您确定你需要确保所有资产有一个资产ID吗？此操作无法撤消。"
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.ensureAssetIDs();

    if (result.error) {
      notify.error("资产 ID 确认失败。");
      return;
    }

    notify.success(`${result.data.completed} 资产更新成功。`);
  }

  async function ensureImportRefs() {
    const { isCanceled } = await confirm.open(
      "您确定你需要确保所有资产有一个引用吗？此操作无法撤消。"
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.ensureImportRefs();

    if (result.error) {
      notify.error("引用导入失败。");
      return;
    }

    notify.success(`${result.data.completed} 资产更新成功。`);
  }

  async function resetItemDateTimes() {
    const { isCanceled } = await confirm.open(
      "您确定你需要删除所有的日期时间吗？此操作无法撤消。"
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.resetItemDateTimes();

    if (result.error) {
      notify.error("日期时间重置失败。");
      return;
    }

    notify.success(`${result.data.completed} 资产更新成功。`);
  }

  async function setPrimaryPhotos() {
    const { isCanceled } = await confirm.open(
      "您确定你需要设置一个首图吗？此操作无法撤消。"
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.setPrimaryPhotos();

    if (result.error) {
      notify.error("首图设置失败。");
      return;
    }

    notify.success(`${result.data.completed} 资产更新成功。`);
  }
</script>

<style scoped></style>
