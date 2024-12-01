if (!istable(query)) then
	require("query")
end

query.EnableInfoDetour(true)

print("Detour enabled")
hook.Add("A2S_INFO", "reply", function(ip, port, info)
	
    info.players = 125
    info.map = 'kreml'
	
    return info
end)

hook.Add("A2S_PLAYER", "reply", function(ip, port, info)
    print("A2S_PLAYER from", ip, port)
	
    return {
        {name = "putin", score = 10000, time = 1},
		{name = "biden", score = 1000, time = 2},
		{name = "zelensky", score = 1, time = 3},
    }
end)
