import {createRouter, createWebHashHistory} from "vue-router";
import StartPageNew from "../pages/StartPageNew.vue";
import DownloadsPageNew from "../pages/DownloadsPageNew.vue";
import TutorialsPage from "../pages/TutorialsPage.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: StartPageNew
    },
    {
        path: "/downloads",
        name: "Downloads",
        component: DownloadsPageNew
    },
    {
        path: "/tutorials",
        name: "Tutorials",
        component: TutorialsPage
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        // always scroll to top
        return {top: 0};
    },
});

export default router;
