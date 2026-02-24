function fileManager() {
  return {
    files: [],

    async fetchFiles() {
      try {
        const response = await fetch("/api/v1/file/list");
        const data = await response.json();

        this.files = data.map((file) => ({ ...file, selected: false }));
      } catch (error) {
        console.error("Could not get file list:", error);
      }
    },

    toggleSelection: function (index) {
      this.files[index].selected = !this.files[index].selected;
    },
  };
}
var cingil = function (element) {
  element.classList.toggle("bg-red-500");
};
