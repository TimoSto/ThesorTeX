import {defineStore} from "pinia";

export const useErrorSuccessStore = defineStore({
    id: "error-success",
    state: () => ({
        valid: true,
        message: ""
    }),

    actions: {
        clear() {
            this.message = "";
            this.valid = true;
        },
        setMessage(valid: boolean, message: string) {
            this.clear();
            setTimeout(() => {
                this.message = "";
                this.valid = valid;
                this.message = message;
            });
        }
    }
});
