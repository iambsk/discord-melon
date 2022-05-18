const { SlashCommandBuilder } = require('@discordjs/builders');
const { REST } = require('@discordjs/rest');
const { Routes } = require('discord-api-types/v9');

const { Client, Collection, Intents, MessageEmbed } = require('discord.js');
const { Discord } = require('discord.js');

const { token, prefix } = require('./config.json');

const allIntents = new Intents//(32767);
const client = new Client({ intents: [allIntents] });



client.commands = new Collection();
client.events = new Collection();
client.on('interactionCreate', async interaction => {
	if (!interaction.isCommand()) return;

	const command = client.commands.get(interaction.commandName);

	if (!command) return;

	try {
		await command.execute(interaction);
	} catch (error) {
		console.error(error);
		await interaction.reply({ content: 'There was an error while executing this command!', ephemeral: true });
	}
});

['command_handler', 'event_handler'].forEach(handler => {
    require(`./handlers/${handler}`)(client, Discord)
})










/*
client.on('messageCreate', message => {
    if (!message.content.startsWith(prefix) || message.author.bot) return;

    const args = message.content.slice(prefix.length).split(/ +/);
    const command = args.shift().toLowerCase();
    switch (command) {
        case 'ping':
            client.commands.get('ping').execute(message, args);
            break;
        case 'embed':
            client.commands.get('embed').execute(message, args, MessageEmbed);
            break;
        default:
            message.reply('Please enter a valid command.');
            break;

    }
});
*/

client.login(token);
