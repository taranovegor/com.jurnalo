<template>
    <div class="grid">
        <div class="col-2">
            <Menu :model="menu">
                <template #start>
                    <span class="inline-flex px-3">
                        <h1 class="text-xl">ĵurnalo!</h1>
                    </span>
                </template>
            </Menu>
        </div>
        <div class="col-10">
            <Panel header="Dashboard">
                <DataTable lazy paginator :value="items" :rows="perPage" :totalRecords="totalRows" @page="loadPage" v-model:expandedRows="expanded" dataKey="_id">
                    <Column field="_realtime" header="Time" style="width: 8em">
                        <template #body="slotProps">
                            <div>{{ formatDate(slotProps.data._realtime, 'L') }}</div>
                            <nobr>{{ formatDate(slotProps.data._realtime, 'LTS') }}</nobr>
                        </template>
                    </Column>
                    <Column header="Entry">
                        <template #body="slotProps">
                            <div class="entry__header">
                                <Badge
                                    v-if="slotProps.data.PRIORITY !== undefined"
                                    :value="priorityToText(slotProps.data.PRIORITY)"
                                    severity="contrast"
                                    class="entry__priority" />
                                <span class="entry__identifier ml-2 vertical-align-middle">{{ slotProps.data._EXE ?? slotProps.data.SYSLOG_IDENTIFIER ?? 'N/A' }}</span>
                            </div>
                            <div class="entry__body mt-1">
                                <span class="entry__message">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab animi aut debitis eos exercitationem fuga hic impedit ipsa numquam, odio pariatur quod recusandae sequi tenetur unde veniam voluptates. Ad aliquam asperiores deleniti eos hic itaque nulla obcaecati reiciendis sed voluptate. Architecto consectetur debitis fugiat minus mollitia nemo quae tempora voluptatem.</span>

                            </div>
                        </template>
                    </Column>
                    <Column expander style="width: 3em" />
                    <template #expansion="slotProps">
                        <DataTable :value="Object.entries(slotProps.data)">
                            <Column field="0" header="Field" style="width: 1em"/>
                            <Column field="1" header="Value"/>
                        </DataTable>
                    </template>
                </DataTable>
            </Panel>
        </div>
    </div>
</template>


<script>
import axios from 'axios'
import moment from 'moment'

export default {
    data() {
        return {
            expanded: [],
            items: [],
            perPage: 12,
            totalRows: 0,
            menu: [
                {
                    separator: true
                },
                {
                    label: 'App',
                    items: [
                        {
                            label: 'Dashboard',
                            icon: 'pi pi-desktop',
                        },
                    ]
                },
            ],
        }
    },
    mounted() {
        this.loadPage()
    },
    methods: {
        loadPage(e) {
            axios.get('http://slave.dot:5341/api/entry', {params: {offset: e?.first ?? 0, limit: this.perPage}}).then(response => {
                this.items = response.data.items
                this.totalRows = response.data.totalCount
                this.expanded = []
            })
        },
        formatDate(string, format) {
            return moment(string).format(format)
        },
        priorityToText(int) {
            return ["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"][int]
        },
    },
}
</script>

<style scoped>
.entry__message {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    line-clamp: 1;
    -webkit-box-orient: vertical;
}
</style>
