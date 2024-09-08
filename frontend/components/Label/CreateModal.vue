<template>
  <BaseModal v-model="modal">
    <template #title> 创建标签 </template>
    <form @submit.prevent="create()">
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="标签名字"
      />
      <FormTextArea v-model="form.description" label="标签描述" />
      <div class="modal-action">
        <div class="flex justify-center">
          <BaseButton class="rounded-r-none" :loading="loading" type="submit"> 创建 </BaseButton>
          <div class="dropdown dropdown-top">
            <label tabindex="0" class="btn rounded-l-none rounded-r-xl">
              <MdiChevronDown class="h-5 w-5" />
            </label>
            <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-64 right-0">
              <li>
                <button type="button" @click="create(false)">创建和添加另一个</button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </form>
    <p class="text-sm text-center mt-4">
      使用 <kbd class="kbd kbd-xs">Shift</kbd> + <kbd class="kbd kbd-xs"> Enter </kbd> 来创建和添加另一个
    </p>
  </BaseModal>
</template>

<script setup lang="ts">
  import MdiChevronDown from "~icons/mdi/chevron-down";
  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const modal = useVModel(props, "modelValue");
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    name: "",
    description: "",
    color: "", // Future!
  });

  function reset() {
    form.name = "";
    form.description = "";
    form.color = "";
    focused.value = false;
    loading.value = false;
  }

  whenever(
    () => modal.value,
    () => {
      focused.value = true;
    }
  );

  const api = useUserApi();
  const toast = useNotifier();

  const { shift } = useMagicKeys();

  async function create(close = true) {
    if (shift.value) {
      close = false;
    }

    const { error, data } = await api.labels.create(form);
    if (error) {
      toast.error("无法创建标签");
      loading.value = false;
      return;
    }

    toast.success("标签已创建");
    reset();

    if (close) {
      modal.value = false;
      navigateTo(`/label/${data.id}`);
    }
  }
</script>
