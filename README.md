# OnlineStatus

Keep code lean by not having to register 'failure' after failed calls.

Instead use a 'maintenance' loop to update the status

```
    for index := range disks {
        if !online.IsOnline(index) {
            continue // skip this disk or node
        }
        
        // perform operation on 'index'
    }    
```

Advantages 
- Simplifies calling code / don't 'clutter' unnecessarily

Disadvantages
- Maybe a delay between first detecting it and it taking effect