#!/usr/bin/env node
import {execSync} from "child_process";
import {lstatSync, readlinkSync, symlinkSync, unlinkSync} from "fs";
import {join, dirname, resolve} from "path";
import {fileURLToPath} from "url";

const configFolder = resolve(join(dirname(fileURLToPath(import.meta.url)), ".."));
const viteConfig = join(configFolder, "vite.config.ts");
const eslintConfig = join(configFolder, ".eslintrc");
const eslintIgnore = join(configFolder, ".eslintignore");
const prettierIgnore = join(configFolder, ".prettierignore");

const linkFile = (source: string, target: string): void => {
  try {
    const stat = lstatSync(target);
    if (!stat.isSymbolicLink()) {
      console.log(`you are using a custom ${target}, aborting`);
      return;
    }
    const linkTo = readlinkSync(target);
    if (linkTo === source) {
      console.log(`${target} already setup, skipping`);
      return;
    }
    console.log(`recreating link to ${target}`);
    unlinkSync(target);
  } catch {
    // eslint-ignore-line no-empty
  }
  symlinkSync(source, target);
};

const ensureSetup = (): void => {
  linkFile(resolve("."), join(configFolder, "user"));
  linkFile(join(configFolder, ".prettierrc.json"), ".prettierrc.json");
  linkFile(join(configFolder, "tsconfig.user.json"), "tsconfig.json");
  console.log("setup successfully");
};

const commands = {
  setup: ensureSetup,
  dev: `vite -c ${viteConfig}`,
  build: `tsc && vite -c ${viteConfig} build`,
  types: `tsc --noEmit`,
  lint: `eslint -c ${eslintConfig} --ignore-path ${eslintIgnore} . --ext .ts --ext tsx`,
  format: `prettier --ignore-path ${prettierIgnore} --check .`,
  "format-fix": `prettier --ignore-path ${prettierIgnore} --write .`,
  clean: "rm -rf dist",
};

const main = (): void => {
  const availableCommands = Object.keys(commands);
  if (process.argv.length < 3 || !availableCommands.includes(process.argv[2])) {
    throw Error(`missing command to execute: ${availableCommands.join(", ")}`);
  }
  const commandName = process.argv[2] as keyof typeof commands;
  const command = commands[commandName];
  if (typeof command !== "string") {
    command();
    return;
  }

  const rest = process.argv.slice(3).join(" ");
  const fullCommand = `${command} ${rest}`;

  console.log(`running '${fullCommand}'`);
  try {
    execSync(`npx -c '${fullCommand}'`, {stdio: "inherit"});
  } catch {
    process.exit(1);
  }
};

main();
