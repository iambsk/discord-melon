const { SlashCommandBuilder } = require('@discordjs/builders');
module.exports = {
    name: 'embed',
    description: 'Testing Embeds',
    execute(client, message , args, Discord){
        const newEmbed = new Discord.MessageEmbed()
        .setColor('#304281')
        .setTitle('Memes')
        // .setURL('https://cdn.discordapp.com/attachments/976608181105066074/976617935797903360/unknown-101.png?size=4096')
        .setDescription('This is an embed for memes.')
        .addFields(
            {name: 'Meme1', value: 'boobies'},
            {name: 'Meme2', value: 'boobies2'}
        )
        .setImage('https://cdn.discordapp.com/attachments/976608181105066074/976617935797903360/unknown-101.png?size=4096');
        // .setFooter('Make sure you boobs');

        message.channel.send({ embeds: [newEmbed] });
    }
    
}