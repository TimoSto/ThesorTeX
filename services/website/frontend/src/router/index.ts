import {createRouter, createWebHashHistory} from "vue-router";
import StartPage from "../pages/StartPage.vue";
import DownloadsPage from "../pages/DownloadsPage.vue";
import TutorialsPage from "../pages/TutorialsPage.vue";

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
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

export default router;
