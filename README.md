# SkriptMC-Parser
SkriptMC-Parser is currently a prototype in the early stages of development of a system that allows the Skript community to test their scripts via a public API for potential errors or warnings. This is a quick and easy way to check your scripts without having to set up a Spigot server on your environment.

## Related projects

Many repositories will integrate it, as part of its operation or as part of using its API. First, the [SkriptMC-Parser-Engine](https://github.com/Romitou/SkriptMC-Parser-Engine) *(Java)* repository is essential to the operation of [SkriptMC-Parser](https://github.com/Romitou/SkriptMC-Parser) *(Go)* as it contains the source code of the parser environment, distributing private Docker images ready to be executed.

It might be possible to use this project to integrate it with Skript-MC's Discord bot, namely [Swan](https://github.com/Skript-MC/Swan) *(TypeScript)*. Users will be able to check the validity of their script with the bot and consequently with SkriptMC-Parser.

Finally, we are surely planning to create an online parser space for the whole community to benefit from and to have a conducive development environment, via the [SkriptMC-Parser-Web](https://github.com/Romitou/SkriptMC-Parser-Web) *(Vue)* repository.

## Perspectives

I think that a parser is not enough for a beginner user as well as for advanced users. I would like to add one more element that will make the difference, namely the execution of the code, in order to have a result of the script. Real-time modification and viewing of variables while the code is running, persistent saving of those variables, sharing and multi-development, one-click addon capability, ultra-fast deployment of parser runtime environments, and much more.