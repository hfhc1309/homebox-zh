<script setup lang="ts">
  import type { LocationSummary, LocationUpdate } from "~~/lib/api/types/data-contracts";
  import { useLocationStore } from "~~/stores/locations";
  import MdiPackageVariant from "~icons/mdi/package-variant";
  import MdiPencil from "~icons/mdi/pencil";
  import MdiDelete from "~icons/mdi/delete";

  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const locationId = computed<string>(() => route.params.id as string);

  const { data: location } = useAsyncData(locationId.value, async () => {
    const { data, error } = await api.locations.get(locationId.value);
    if (error) {
      toast.error("无法加载位置");
      navigateTo("/home");
      return;
    }

    if (data.parent) {
      parent.value = locations.value.find(l => l.id === data.parent.id);
    }

    return data;
  });

  const confirm = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await confirm.open(
      "您确定要删除此位置及其所有项目吗？此操作无法撤消。"
    );
    if (isCanceled) {
      return;
    }

    const { error } = await api.locations.delete(locationId.value);
    if (error) {
      toast.error("位置无法删除");
      return;
    }

    toast.success("位置已删除");
    navigateTo("/home");
  }

  const updateModal = ref(false);
  const updating = ref(false);
  const updateData = reactive<LocationUpdate>({
    id: locationId.value,
    name: "",
    description: "",
    parentId: null,
  });

  function openUpdate() {
    updateData.name = location.value?.name || "";
    updateData.description = location.value?.description || "";
    updateModal.value = true;
  }

  async function update() {
    updating.value = true;
    updateData.parentId = parent.value?.id || null;
    const { error, data } = await api.locations.update(locationId.value, updateData);

    if (error) {
      updating.value = false;
      toast.error("位置无法更新");
      return;
    }

    toast.success("位置已更新");
    location.value = data;
    updateModal.value = false;
    updating.value = false;
  }

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.allLocations);

  const parent = ref<LocationSummary | any>({});

  const items = computedAsync(async () => {
    if (!location.value) {
      return [];
    }

    const resp = await api.items.getAll({
      locations: [location.value.id],
    });

    if (resp.error) {
      toast.error("无法加载资产");
      return [];
    }

    return resp.data.items;
  });
</script>

<template>
  <div>
    <!-- Update Dialog -->
    <BaseModal v-model="updateModal">
      <template #title> 更新位置 </template>
      <form v-if="location" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="位置名称" />
        <FormTextArea v-model="updateData.description" label="位置描述" />
        <LocationSelector v-model="parent" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> 更新 </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseContainer v-if="location">
      <div class="bg-base-100 rounded p-3">
        <header class="mb-2">
          <div class="flex flex-wrap items-end gap-2">
            <div class="avatar placeholder mb-auto">
              <div class="bg-neutral-focus text-neutral-content rounded-full w-12">
                <MdiPackageVariant name="mdi-package-variant" class="h-7 w-7" />
              </div>
            </div>
            <div>
              <div v-if="location?.parent" class="text-sm breadcrumbs pt-0 pb-0">
                <ul class="text-base-content/70">
                  <li>
                    <NuxtLink :to="`/location/${location.parent.id}`"> {{ location.parent.name }}</NuxtLink>
                  </li>
                  <li>{{ location.name }}</li>
                </ul>
              </div>
              <h1 class="text-2xl pb-1">
                {{ location ? location.name : "" }}
              </h1>
              <div class="flex gap-1 flex-wrap text-xs">
                <div>
                  创建
                  <DateTime :date="location?.createdAt" />
                </div>
              </div>
            </div>
            <div class="ml-auto mt-2 flex flex-wrap items-center justify-between gap-3">
              <div class="btn-group">
                <PageQRCode class="dropdown-left" />
                <BaseButton size="sm" @click="openUpdate">
                  <MdiPencil class="mr-1" name="mdi-pencil" />
                  编辑
                </BaseButton>
              </div>
              <BaseButton class="btn btn-sm" @click="confirmDelete()">
                <MdiDelete name="mdi-delete" class="mr-2" />
                删除
              </BaseButton>
            </div>
          </div>
        </header>
        <div class="divider my-0 mb-1"></div>
        <Markdown v-if="location && location.description" class="text-base" :source="location.description"> </Markdown>
      </div>
      <section v-if="location && items">
        <ItemViewSelectable :items="items" />
      </section>

      <section v-if="location && location.children.length > 0" class="mt-6">
        <BaseSectionHeader class="mb-5"> Child Locations </BaseSectionHeader>
        <div class="grid gap-2 grid-cols-1 sm:grid-cols-3">
          <LocationCard v-for="item in location.children" :key="item.id" :location="item" />
        </div>
      </section>
    </BaseContainer>
  </div>
</template>
