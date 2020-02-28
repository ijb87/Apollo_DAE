<template>
    <Page>
        <ActionBar title="Home" />
        <ScrollView>
            <StackLayout class="home-panel">
                <Image
                    src="https://play.nativescript.org/dist/assets/img/NativeScript_logo.png" />
                <Button class="btn btn-primary" text="Button" @tap="onButtonTap" />
            </StackLayout>
        </ScrollView>
    </Page>
</template>

<script>
    export default {
        methods: {
            onButtonTap() {
                console.log("Button was pressed");
            }
        },

        data() {
            return {};
        }
    };
</script>

<style scoped>
    .home-panel {
        vertical-align: center;
        font-size: 20;
        margin: 15;
    }

    .description-label {
        margin-bottom: 15;
    }
</style>