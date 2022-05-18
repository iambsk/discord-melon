const { SlashCommandBuilder } = require('@discordjs/builders');
module.exports = {
    name: 'ping',
    description: 'This is a ping command',
    execute(client, message , args, Discord){
        console.log('Melon bot is online!');
// ADMINISTRATOR
        // if(!message.member.roles.cache.has('976614921913630740')){
        // if(message.member.roles.cache.has(r => r.name === "admin")){
        if(message.member.permissions.has("ADMINISTRATOR")){
            message.reply('You have admin permissions.\npong!');
            
        } else{
            message.reply('pong!');
        }
        
    }
}