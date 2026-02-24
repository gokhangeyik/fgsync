const plugin = require("tailwindcss/plugin");

module.exports = {
  darkMode: "class", // .dark sınıfını kullanarak dark mode etkinleştirme
  content: ["./src/**/*.{js,ts,jsx,tsx,html}"],
  theme: {
    extend: {
      fontSize: {
        vsm: "0.825rem",
      },
      colors: {
        fujiWhite: "#dcd7ba",
        oldWhite: "#c8c093",
        gokkoWhite: "#c8c7c2",
        sumiInk0: "#16161d",
        sumiInk1: "#1f1f28",
        sumiInk2: "#2a2a37",
        sumiInk3: "#363646",
        sumiInk4: "#54546d",
        waveBlue1: "#223249",
        waveBlue2: "#2d4f67",
        winterGreen: "#2b3328",
        winterYellow: "#49443c",
        winterRed: "#43242b",
        winterBlue: "#252535",
        autumnGreen: "#76946a",
        autumnRed: "#c34043",
        autumnYellow: "#dca561",
        samuraiRed: "#e82424",
        roninYellow: "#ff9e3b",
        waveAqua1: "#6a9589",
        dragonBlue: "#658594",
        fujiGray: "#727169",
        springViolet1: "#938aa9",
        springViolet2: "#9cabca",
        oniViolet: "#957fb8",
        crystalBlue: "#7e9cd8",
        springBlue: "#7fb4ca",
        lightBlue: "#a3d4d5",
        waveAqua2: "#7aa89f",
        springGreen: "#98bb6c",
        boatYellow1: "#938056",
        boatYellow2: "#c0a36e",
        carpYellow: "#e6c384",
        sakuraPink: "#d27e99",
        waveRed: "#e46876",
        peachRed: "#ff5d62",
        surimiOrange: "#ffa066",
      },
      borderRadius: {
        "tl-lg": "0.5rem",
        "br-lg": "0.5rem",
        "bl-sm": "0.25rem",
        "tr-sm": "0.25rem",
      },
      transitionDuration: {
        200: "200ms",
      },
    },
  },
  plugins: [
    plugin(function ({ addVariant }) {
      addVariant("dark", "&:where(.dark, .dark *)");
    }),
  ],
};
