<script setup lang="ts">
  import MdiPackageVariant from "~icons/mdi/package-variant";
  import MdiPencil from "~icons/mdi/pencil";
  import MdiDelete from "~icons/mdi/delete";

  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const labelId = computed<string>(() => route.params.id as string);

  const { data: label } = useAsyncData(labelId.value, async () => {
    const { data, error } = await api.labels.get(labelId.value);
    if (error) {
      toast.error("无法加载标签");
      navigateTo("/home");
      return;
    }
    return data;
  });

  const confirm = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await confirm.open(
      "你确定要删除这一个标签吗？这一个操作不能撤回。"
    );

    if (isCanceled) {
      return;
    }

    const { error } = await api.labels.delete(labelId.value);

    if (error) {
      toast.error("标签无法删除");
      return;
    }
    toast.success("标签已删除");
    navigateTo("/home");
  }

  const updateModal = ref(false);
  const updating = ref(false);
  const updateData = reactive({
    name: "",
    description: "",
    color: "",
  });

  function openUpdate() {
    updateData.name = label.value?.name || "";
    updateData.description = label.value?.description || "";
    updateModal.value = true;
  }

  async function update() {
    updating.value = true;
    const { error, data } = await api.labels.update(labelId.value, updateData);

    if (error) {
      updating.value = false;
      toast.error("标签无法更新");
      return;
    }

    toast.success("标签已更新");
    label.value = data;
    updateModal.value = false;
    updating.value = false;
  }

  const items = computedAsync(async () => {
    if (!label.value) {
      return [];
    }

    const resp = await api.items.getAll({
      labels: [label.value.id],
    });

    if (resp.error) {
      toast.error("无法加载资产");
      return [];
    }

    return resp.data.items;
  });
</script>

<template>
  <BaseContainer>
    <BaseModal v-model="updateModal">
      <template #title> Update Label </template>
      <form v-if="label" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="Label Name" />
        <FormTextArea v-model="updateData.description" label="Label Description" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseContainer v-if="label">
      <div class="bg-base-100 rounded p-3">
        <header class="mb-2">
          <div class="flex flex-wrap items-end gap-2">
            <div class="avatar placeholder mb-auto">
              <div class="bg-neutral-focus text-neutral-content rounded-full w-12">
                <MdiPackageVariant class="h-7 w-7" />
              </div>
            </div>
            <div>
              <h1 class="text-2xl pb-1">
                {{ label ? label.name : "" }}
              </h1>
              <div class="flex gap-1 flex-wrap text-xs">
                <div>
                  创建
                  <DateTime :date="label?.createdAt" />
                </div>
              </div>
            </div>
            <div class="ml-auto mt-2 flex flex-wrap items-center justify-between gap-3">
              <div class="btn-group">
                <PageQRCode class="dropdown-left" />
                <BaseButton size="sm" @click="openUpdate">
                  <MdiPencil class="mr-1" />
                  编辑
                </BaseButton>
              </div>
              <BaseButton class="btn btn-sm" @click="confirmDelete()">
                <MdiDelete class="mr-2" />
                删除
              </BaseButton>
            </div>
          </div>
        </header>
        <div class="divider my-0 mb-1"></div>
        <Markdown v-if="label && label.description" class="text-base" :source="label.description"> </Markdown>
      </div>
      <section v-if="label && items">
        <ItemViewSelectable :items="items" />
      </section>
    </BaseContainer>
  </BaseContainer>
</template>
