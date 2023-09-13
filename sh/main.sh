#!/usr/bin/awk -f
{
    s = 0;
    for(i = 3; i <= NF; i++) {
        if ($i == 0) { next }
        s += $i * $2;
    }

    for (i in names) {
        if (names[i] == $1) {
            all[i] += s;
            noZeros[i] ++;
            next
        }
    }

    l = length(names)
    names[l] = $1
    all[l] = s;
    noZeros[l] = 1;
}
END{
    l = length(names)
    for (i=0 ; i<l ; i++) {
        printf "%s %d %.2f\n", names[i], noZeros[i], all[i]/100000;
    }
}