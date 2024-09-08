import type { ComputedRef, Ref } from "vue";
import { scorePassword } from "~~/lib/passwords";

export interface PasswordScore {
  score: ComputedRef<number>;
  message: ComputedRef<string>;
  isValid: ComputedRef<boolean>;
}

export function usePasswordScore(pw: Ref<string>, min = 30): PasswordScore {
  const score = computed(() => {
    return scorePassword(pw.value) || 0;
  });

  const message = computed(() => {
    if (score.value < 20) {
      return "非常弱";
    } else if (score.value < 40) {
      return "弱";
    } else if (score.value < 60) {
      return "好";
    } else if (score.value < 80) {
      return "强";
    }
    return "非常强";
  });

  const isValid = computed(() => {
    return score.value >= min;
  });

  return {
    score,
    isValid,
    message,
  };
}
