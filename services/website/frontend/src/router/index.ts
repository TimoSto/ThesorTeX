import {createRouter, createWebHashHistory} from "vue-router";
import StartPage from "../pages/StartPage.vue";
import DownloadsPage from "../pages/DownloadsPage.vue";
import TutorialsPage from "../pages/TutorialsPage.vue";
import ImprintPage from "../pages/ImprintPage.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: StartPage
    },
    {
        path: "/downloads",
        name: "Downloads",
        component: DownloadsPage
    },
    {
        path: "/tutorials",
        name: "Tutorials",
        component: TutorialsPage
    },
    {
        path: "/imprint",
        name: "Imprint",
        component: ImprintPage
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
